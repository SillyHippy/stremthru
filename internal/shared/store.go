package shared

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/SillyHippy/stremthru/core"
	"github.com/SillyHippy/stremthru/internal/cache"
	"github.com/SillyHippy/stremthru/internal/config"
	"github.com/SillyHippy/stremthru/internal/context"
	"github.com/SillyHippy/stremthru/store"
	"github.com/SillyHippy/stremthru/store/alldebrid"
	"github.com/SillyHippy/stremthru/store/debridlink"
	"github.com/SillyHippy/stremthru/store/easydebrid"
	"github.com/SillyHippy/stremthru/store/offcloud"
	"github.com/SillyHippy/stremthru/store/premiumize"
	"github.com/SillyHippy/stremthru/store/realdebrid"
	"github.com/SillyHippy/stremthru/store/torbox"
	"github.com/golang-jwt/jwt/v5"
)

var adStore = alldebrid.NewStoreClient(&alldebrid.StoreClientConfig{
	HTTPClient: GetHTTPClient(config.StoreTunnel.IsEnabledForAPI("alldebrid")),
})
var dlStore = debridlink.NewStoreClient(&debridlink.StoreClientConfig{
	HTTPClient: GetHTTPClient(config.StoreTunnel.IsEnabledForAPI("debridlink")),
})
var edStore = easydebrid.NewStoreClient(&easydebrid.StoreClientConfig{
	HTTPClient: GetHTTPClient(config.StoreTunnel.IsEnabledForAPI("easydebrid")),
})
var pmStore = premiumize.NewStoreClient(&premiumize.StoreClientConfig{
	HTTPClient: GetHTTPClient(config.StoreTunnel.IsEnabledForAPI("premiumize")),
})
var ocStore = offcloud.NewStoreClient(&offcloud.StoreClientConfig{
	HTTPClient: GetHTTPClient(config.StoreTunnel.IsEnabledForAPI("offcloud")),
})
var rdStore = realdebrid.NewStoreClient(&realdebrid.StoreClientConfig{
	HTTPClient: GetHTTPClient(config.StoreTunnel.IsEnabledForAPI("realdebrid")),
	UserAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
})
var tbStore = torbox.NewStoreClient(&torbox.StoreClientConfig{
	HTTPClient: GetHTTPClient(config.StoreTunnel.IsEnabledForAPI("torbox")),
})

func GetStore(name string) store.Store {
	switch store.StoreName(name) {
	case store.StoreNameAlldebrid:
		return adStore
	case store.StoreNameDebridLink:
		return dlStore
	case store.StoreNameEasyDebrid:
		return edStore
	case store.StoreNameOffcloud:
		return ocStore
	case store.StoreNamePremiumize:
		return pmStore
	case store.StoreNameRealDebrid:
		return rdStore
	case store.StoreNameTorBox:
		return tbStore
	default:
		return nil
	}
}

type proxyLinkTokenData struct {
	EncLink    string `json:"enc_link"`
	EncFormat  string `json:"enc_format"`
	SkipTunnel bool   `json:"notun"`
}

func CreateProxyLink(r *http.Request, ctx *context.RequestContext, link string) (string, error) {
	if !ctx.IsProxyAuthorized || ctx.StoreAuthToken != config.StoreAuthToken.GetToken(ctx.ProxyAuthUser, string(ctx.Store.GetName())) {
		return link, nil
	}

	encryptedLink, err := core.Encrypt(ctx.ProxyAuthPassword, link)
	if err != nil {
		return "", err
	}

	proxyLink := ExtractRequestBaseURL(r).JoinPath("/v0/store/link/access")
	useTunnel := config.StoreTunnel.IsEnabledForStream(string(ctx.Store.GetName()))

	token, err := core.CreateJWT(ctx.ProxyAuthPassword, core.JWTClaims[proxyLinkTokenData]{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "stremthru",
			Subject:   ctx.ProxyAuthUser,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
		Data: &proxyLinkTokenData{
			EncLink:    encryptedLink,
			EncFormat:  core.EncryptionFormat,
			SkipTunnel: !useTunnel,
		},
	})

	if err != nil {
		return "", err
	}

	proxyLink = proxyLink.JoinPath(token)

	if filename := filepath.Base(link); filename != "" {
		proxyLink = proxyLink.JoinPath(filename)
	}

	return proxyLink.String(), nil
}

func GenerateStremThruLink(r *http.Request, ctx *context.RequestContext, link string) (*store.GenerateLinkData, error) {
	params := &store.GenerateLinkParams{}
	params.APIKey = ctx.StoreAuthToken
	params.Link = link
	if ctx.ClientIP != "" {
		params.ClientIP = ctx.ClientIP
	}

	data, err := ctx.Store.GenerateLink(params)
	if err != nil {
		return nil, err
	}

	proxyLink, err := CreateProxyLink(r, ctx, data.Link)
	if err != nil {
		return nil, err
	}

	data.Link = proxyLink

	return data, nil
}

type proxyLink struct {
	Value string
	NoTun bool
}

var proxyLinkTokenCache = func() cache.Cache[proxyLink] {
	return cache.NewCache[proxyLink](&cache.CacheConfig{
		Name:     "store:proxyLinkToken",
		Lifetime: 30 * time.Minute,
	})
}()

func getUserSecretFromJWT(t *jwt.Token) (string, []byte, error) {
	username, err := t.Claims.GetSubject()
	if err != nil {
		return "", nil, err
	}
	password := config.ProxyAuthPassword.GetPassword(username)
	return password, []byte(password), nil
}

func UnwrapProxyLinkToken(encodedToken string) (string, bool, error) {
	proxyLink := &proxyLink{}
	if found := proxyLinkTokenCache.Get(encodedToken, proxyLink); found {
		return proxyLink.Value, proxyLink.NoTun, nil
	}

	claims := &core.JWTClaims[proxyLinkTokenData]{}
	token, err := core.ParseJWT(func(t *jwt.Token) (interface{}, error) {
		_, secret, err := getUserSecretFromJWT(t)
		return secret, err
	}, encodedToken, claims)

	if err != nil {
		return "", false, err
	}

	secret, _, err := getUserSecretFromJWT(token)
	if err != nil {
		return "", false, err
	}

	proxyLink.NoTun = !claims.Data.SkipTunnel
	proxyLink.Value, err = core.Decrypt(secret, claims.Data.EncLink)
	if err != nil {
		return "", false, err
	}

	proxyLinkTokenCache.Add(encodedToken, *proxyLink)

	return proxyLink.Value, proxyLink.NoTun, nil
}
