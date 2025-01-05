package config

import (
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/SillyHippy/stremthru/core"
	"github.com/SillyHippy/stremthru/store"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && len(value) > 0 {
		return value
	}
	return defaultValue
}

type StoreAuthTokenMap map[string]map[string]string

func (m StoreAuthTokenMap) GetToken(user, store string) string {
	if um, ok := m[user]; ok {
		if token, ok := um[store]; ok {
			return token
		}
	}
	if user != "*" {
		return m.GetToken("*", store)
	}
	return ""
}

func (m StoreAuthTokenMap) setToken(user, store, token string) {
	if _, ok := m[user]; !ok {
		m[user] = make(map[string]string)
	}
	m[user][store] = token
}

func (m StoreAuthTokenMap) GetPreferredStore(user string) string {
	store := m.GetToken(user, "*")
	if store == "" {
		store = m.GetToken("*", "*")
	}
	return store
}

func (m StoreAuthTokenMap) ListStores(user string) []string {
	stores := []string{}
	if um, ok := m[user]; ok {
		for store := range um {
			if store != "*" {
				stores = append(stores, store)
			}
		}
	}
	return stores
}

func (m StoreAuthTokenMap) setPreferredStore(user, store string) {
	if m.GetPreferredStore(user) == "" {
		m.setToken(user, "*", store)
	}
}

type ProxyAuthPasswordMap map[string]string

func (m ProxyAuthPasswordMap) GetPassword(userName string) string {
	if token, ok := m[userName]; ok {
		return token
	}
	return ""
}

const (
	StremioAddonSidekick string = "sidekick"
	StremioAddonStore    string = "store"
	StremioAddonWrap     string = "wrap"
)

var stremioAddons = []string{StremioAddonSidekick, StremioAddonStore, StremioAddonWrap}

type StremioAddonConfig struct {
	enabled []string
}

func (sa StremioAddonConfig) IsEnabled(name string) bool {
	if len(sa.enabled) == 0 {
		return true
	}

	for _, addon := range sa.enabled {
		if addon == name {
			return true
		}
	}
	return false
}

type StoreTunnelConfig struct {
	api    bool
	stream bool
}

type StoreTunnelConfigMap map[string]StoreTunnelConfig

func (stc StoreTunnelConfigMap) IsEnabledForAPI(name string) bool {
	if c, ok := stc[name]; ok {
		return c.api
	}
	if name != "*" {
		return stc.IsEnabledForAPI("*")
	}
	return true
}

func (stc StoreTunnelConfigMap) IsEnabledForStream(name string) bool {
	if c, ok := stc[name]; ok {
		return c.stream
	}
	if name != "*" {
		return stc.IsEnabledForStream("*")
	}
	return true
}

type Config struct {
	Port               string
	StoreAuthToken     StoreAuthTokenMap
	ProxyAuthPassword  ProxyAuthPasswordMap
	ProxyStreamEnabled bool
	BuddyURL           string
	HasBuddy           bool
	PeerURL            string
	PeerAuthToken      string
	HasPeer            bool
	RedisURI           string
	DatabaseURI        string
	StremioAddon       StremioAddonConfig
	Version            string
	LandingPage        string
	ServerStartTime    time.Time
	StoreTunnel        StoreTunnelConfigMap
}

func parseUri(uri string) (parsedUrl, parsedToken string) {
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatalf("invalid uri: %s", uri)
	}
	if password, ok := u.User.Password(); ok {
		parsedToken = password
	} else {
		parsedToken = u.User.Username()
	}
	u.User = nil
	parsedUrl = strings.TrimSpace(u.String())
	return
}

var config = func() Config {
	if value := getEnv("STREMTHRU_HTTP_PROXY", ""); len(value) > 0 {
		if err := os.Setenv("HTTP_PROXY", value); err != nil {
			log.Fatal("failed to set http proxy")
		}
	}

	if value := getEnv("STREMTHRU_HTTPS_PROXY", ""); len(value) > 0 {
		if err := os.Setenv("HTTPS_PROXY", value); err != nil {
			log.Fatal("failed to set https proxy")
		}
	}

	proxyAuthCredList := strings.FieldsFunc(getEnv("STREMTHRU_PROXY_AUTH", ""), func(c rune) bool {
		return c == ','
	})
	proxyAuthPasswordMap := make(ProxyAuthPasswordMap)
	for _, cred := range proxyAuthCredList {
		if basicAuth, err := core.ParseBasicAuth(cred); err == nil {
			proxyAuthPasswordMap[basicAuth.Username] = basicAuth.Password
		}
	}

	storeAlldebridTokenList := strings.FieldsFunc(getEnv("STREMTHRU_STORE_AUTH", ""), func(c rune) bool {
		return c == ','
	})
	storeAuthTokenMap := make(StoreAuthTokenMap)
	for _, userStoreToken := range storeAlldebridTokenList {
		if user, storeToken, ok := strings.Cut(userStoreToken, ":"); ok {
			if store, token, ok := strings.Cut(storeToken, ":"); ok {
				storeAuthTokenMap.setPreferredStore(user, store)
				storeAuthTokenMap.setToken(user, store, token)
			}
		}
	}

	buddyUrl, _ := parseUri(getEnv("STREMTHRU_BUDDY_URI", ""))
	peerUrl, peerAuthToken := parseUri(getEnv("STREMTHRU_PEER_URI", ""))

	databaseUri := getEnv("STREMTHRU_DATABASE_URI", "sqlite://./data/stremthru.db")

	stremioAddon := StremioAddonConfig{
		enabled: strings.FieldsFunc(strings.TrimSpace(getEnv("STREMTHRU_STREMIO_ADDON", strings.Join(stremioAddons, ","))), func(c rune) bool {
			return c == ','
		}),
	}

	storeTunnelList := strings.FieldsFunc(getEnv("STREMTHRU_STORE_TUNNEL", "*:true"), func(c rune) bool {
		return c == ','
	})

	storeTunnelMap := make(StoreTunnelConfigMap)
	for _, storeTunnel := range storeTunnelList {
		if store, tunnel, ok := strings.Cut(storeTunnel, ":"); ok {
			storeTunnelMap[store] = StoreTunnelConfig{
				api:    tunnel == "true" || tunnel == "api",
				stream: tunnel == "true",
			}
		}
	}

	return Config{
		Port:               getEnv("PORT", getEnv("STREMTHRU_PORT", "8080")), // Prioritize "PORT" for platforms like Vercel
		ProxyAuthPassword:  proxyAuthPasswordMap,
		ProxyStreamEnabled: len(proxyAuthPasswordMap) > 0,
		StoreAuthToken:     storeAuthTokenMap,
		BuddyURL:           buddyUrl,
		HasBuddy:           len(buddyUrl) > 0,
		PeerURL:            peerUrl,
		PeerAuthToken:      peerAuthToken,
		HasPeer:            len(peerUrl) > 0,
		RedisURI:           getEnv("STREMTHRU_REDIS_URI", ""),
		DatabaseURI:        databaseUri,
		StremioAddon:       stremioAddon,
		Version:            "0.28.1", // x-release-please-version
		LandingPage:        getEnv("STREMTHRU_LANDING_PAGE", "{}"),
		ServerStartTime:    time.Now(),
		StoreTunnel:        storeTunnelMap,
	}
}()

var (
	Port               = config.Port
	ProxyAuthPassword  = config.ProxyAuthPassword
	ProxyStreamEnabled = config.ProxyStreamEnabled
	StoreAuthToken     = config.StoreAuthToken
	BuddyURL           = config.BuddyURL
	HasBuddy           = config.HasBuddy
	PeerURL            = config.PeerURL
	PeerAuthToken      = config.PeerAuthToken
	HasPeer            = config.HasPeer
	RedisURI           = config.RedisURI
	DatabaseURI        = config.DatabaseURI
	StremioAddon       = config.StremioAddon
	Version            = config.Version
	LandingPage        = config.LandingPage
	ServerStartTime    = config.ServerStartTime
	StoreTunnel        = config.StoreTunnel
)

func getRedactedURI(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	return u.Redacted(), nil
}

func PrintConfig() {
	l := log.New(os.Stderr, "=", 0)
	l.Println("====== StremThru =======")
	l.Printf(" Time: %v\n", ServerStartTime.Format(time.RFC3339))
	l.Printf(" Version: %v\n", Version)
	l.Printf(" Port: %v\n", Port)
	l.Println("========================")
}
