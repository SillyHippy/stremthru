package endpoint

import (
	"net/http"
	"strings"

	"github.com/SillyHippy/stremthru/internal/buddy"
	"github.com/SillyHippy/stremthru/internal/context"
	"github.com/SillyHippy/stremthru/internal/peer_token"
	"github.com/SillyHippy/stremthru/internal/shared"
	"github.com/SillyHippy/stremthru/internal/store/video"
	"github.com/SillyHippy/stremthru/store"
)

func getUser(ctx *context.RequestContext) (*store.User, error) {
	params := &store.GetUserParams{}
	params.APIKey = ctx.StoreAuthToken
	return ctx.Store.GetUser(params)
}

func handleStoreUser(w http.ResponseWriter, r *http.Request) {
	if !shared.IsMethod(r, http.MethodGet) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	ctx := context.GetRequestContext(r)
	user, err := getUser(ctx)
	SendResponse(w, 200, user, err)
}

type AddMagnetPayload struct {
	Magnet string `json:"magnet"`
}

func checkMagnet(ctx *context.RequestContext, magnets []string, sid string) (*store.CheckMagnetData, error) {
	params := &store.CheckMagnetParams{}
	params.APIKey = ctx.StoreAuthToken
	params.Magnets = magnets
	params.SId = sid
	if ctx.ClientIP != "" {
		params.ClientIP = ctx.ClientIP
	}
	data, err := ctx.Store.CheckMagnet(params)
	if err == nil && data.Items == nil {
		data.Items = []store.CheckMagnetDataItem{}
	}
	return data, err
}

type TrackMagnetPayload struct {
	Hash   string             `json:"hash"`
	Files  []store.MagnetFile `json:"files"`
	IsMiss bool               `json:"is_miss"`
	SId    string             `json:"sid"`
}

type TrackMagnetData struct {
}

func hadleStoreMagnetsTrack(w http.ResponseWriter, r *http.Request) {
	if !shared.IsMethod(r, http.MethodPost) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	ctx := context.GetRequestContext(r)

	isValidToken, err := peer_token.IsValid(ctx.PeerToken)
	if err != nil {
		SendError(w, err)
		return
	}
	if !isValidToken {
		shared.ErrorUnauthorized(r).Send(w)
		return
	}

	payload := &TrackMagnetPayload{}
	if err := shared.ReadRequestBodyJSON(r, payload); err != nil {
		SendError(w, err)
		return
	}

	buddy.TrackMagnet(ctx.Store, payload.Hash, payload.Files, payload.SId, payload.IsMiss, ctx.StoreAuthToken)

	SendResponse(w, 202, &TrackMagnetData{}, nil)
}

func handleStoreMagnetsCheck(w http.ResponseWriter, r *http.Request) {
	if shared.IsMethod(r, http.MethodPost) {
		hadleStoreMagnetsTrack(w, r)
		return
	}

	if !shared.IsMethod(r, http.MethodGet) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	queryParams := r.URL.Query()
	magnet, ok := queryParams["magnet"]
	if !ok {
		shared.ErrorBadRequest(r, "missing magnet").Send(w)
		return
	}

	magnets := []string{}
	for _, m := range magnet {
		magnets = append(magnets, strings.FieldsFunc(m, func(r rune) bool {
			return r == ','
		})...)
	}

	if len(magnets) == 0 {
		shared.ErrorBadRequest(r, "missing magnet").Send(w)
		return
	}

	sid := queryParams.Get("sid")

	ctx := context.GetRequestContext(r)
	data, err := checkMagnet(ctx, magnets, sid)
	if err == nil && data != nil {
		for _, item := range data.Items {
			item.Hash = strings.ToLower(item.Hash)
		}
	}
	SendResponse(w, 200, data, err)
}

func listMagnets(ctx *context.RequestContext, r *http.Request) (*store.ListMagnetsData, error) {
	queryParams := r.URL.Query()
	limit, err := GetQueryInt(queryParams, "limit", 100)
	if err != nil {
		return nil, shared.ErrorBadRequest(r, err.Error())
	}
	if limit > 500 {
		limit = 500
	}
	offset, err := GetQueryInt(queryParams, "offset", 0)
	if err != nil {
		return nil, shared.ErrorBadRequest(r, err.Error())
	}

	params := &store.ListMagnetsParams{
		Limit:  limit,
		Offset: offset,
	}
	params.APIKey = ctx.StoreAuthToken
	data, err := ctx.Store.ListMagnets(params)

	if err == nil && data.Items == nil {
		data.Items = []store.ListMagnetsDataItem{}
	}
	return data, err
}

func handleStoreMagnetsList(w http.ResponseWriter, r *http.Request) {
	if !shared.IsMethod(r, http.MethodGet) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	ctx := context.GetRequestContext(r)
	data, err := listMagnets(ctx, r)
	if err == nil && data != nil {
		for _, item := range data.Items {
			item.Hash = strings.ToLower(item.Hash)
		}
	}
	SendResponse(w, 200, data, err)
}

func addMagnet(ctx *context.RequestContext, magnet string) (*store.AddMagnetData, error) {
	params := &store.AddMagnetParams{}
	params.APIKey = ctx.StoreAuthToken
	params.Magnet = magnet
	if ctx.ClientIP != "" {
		params.ClientIP = ctx.ClientIP
	}
	data, err := ctx.Store.AddMagnet(params)
	if err == nil {
		buddy.TrackMagnet(ctx.Store, data.Hash, data.Files, "*", data.Status != store.MagnetStatusDownloaded, ctx.StoreAuthToken)
	}
	return data, err
}

func handleStoreMagnetAdd(w http.ResponseWriter, r *http.Request) {
	if !shared.IsMethod(r, http.MethodPost) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	payload := &AddMagnetPayload{}
	err := shared.ReadRequestBodyJSON(r, payload)
	if err != nil {
		SendError(w, err)
		return
	}

	ctx := context.GetRequestContext(r)
	data, err := addMagnet(ctx, payload.Magnet)
	if err == nil && data != nil {
		data.Hash = strings.ToLower(data.Hash)
	}
	SendResponse(w, 201, data, err)
}

func handleStoreMagnets(w http.ResponseWriter, r *http.Request) {
	if shared.IsMethod(r, http.MethodGet) {
		handleStoreMagnetsList(w, r)
		return
	}

	if shared.IsMethod(r, http.MethodPost) {
		handleStoreMagnetAdd(w, r)
		return
	}

	shared.ErrorMethodNotAllowed(r).Send(w)
}

func getMagnet(ctx *context.RequestContext, magnetId string) (*store.GetMagnetData, error) {
	params := &store.GetMagnetParams{}
	params.APIKey = ctx.StoreAuthToken
	params.Id = magnetId
	data, err := ctx.Store.GetMagnet(params)
	if err == nil {
		buddy.TrackMagnet(ctx.Store, data.Hash, data.Files, "*", data.Status != store.MagnetStatusDownloaded, ctx.StoreAuthToken)
	}
	return data, err
}

func handleStoreMagnetGet(w http.ResponseWriter, r *http.Request) {
	if !shared.IsMethod(r, http.MethodGet) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	magnetId := r.PathValue("magnetId")
	if magnetId == "" {
		shared.ErrorBadRequest(r, "missing magnetId").Send(w)
		return
	}

	ctx := context.GetRequestContext(r)
	data, err := getMagnet(ctx, magnetId)
	if err == nil && data != nil {
		data.Hash = strings.ToLower(data.Hash)
	}
	SendResponse(w, 200, data, err)
}

func removeMagnet(ctx *context.RequestContext, magnetId string) (*store.RemoveMagnetData, error) {
	params := &store.RemoveMagnetParams{}
	params.APIKey = ctx.StoreAuthToken
	params.Id = magnetId
	return ctx.Store.RemoveMagnet(params)
}

func handleStoreMagnetRemove(w http.ResponseWriter, r *http.Request) {
	if !shared.IsMethod(r, http.MethodDelete) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	magnetId := r.PathValue("magnetId")
	if magnetId == "" {
		shared.ErrorBadRequest(r, "missing magnetId").Send(w)
		return
	}

	ctx := context.GetRequestContext(r)
	data, err := removeMagnet(ctx, magnetId)
	SendResponse(w, 200, data, err)
}

func handleStoreMagnet(w http.ResponseWriter, r *http.Request) {
	if shared.IsMethod(r, http.MethodGet) {
		handleStoreMagnetGet(w, r)
		return
	}

	if shared.IsMethod(r, http.MethodDelete) {
		handleStoreMagnetRemove(w, r)
		return
	}

	shared.ErrorMethodNotAllowed(r).Send(w)
}

type GenerateLinkPayload struct {
	Link string `json:"link"`
}

func handleStoreLinkGenerate(w http.ResponseWriter, r *http.Request) {
	if !shared.IsMethod(r, http.MethodPost) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	payload := &GenerateLinkPayload{}
	err := shared.ReadRequestBodyJSON(r, payload)
	if err != nil {
		SendError(w, err)
		return
	}

	ctx := context.GetRequestContext(r)
	link, err := shared.GenerateStremThruLink(r, ctx, payload.Link)
	SendResponse(w, 200, link, err)
}

func handleStoreLinkAccess(w http.ResponseWriter, r *http.Request) {
	if !shared.IsMethod(r, http.MethodGet) && !shared.IsMethod(r, http.MethodHead) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	encodedToken := r.PathValue("token")
	if encodedToken == "" {
		shared.ErrorBadRequest(r, "missing token").Send(w)
		return
	}

	link, useTunnel, err := shared.UnwrapProxyLinkToken(encodedToken)
	if err != nil {
		SendError(w, err)
		return
	}

	shared.ProxyResponse(w, r, link, useTunnel)
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	if !shared.IsMethod(r, http.MethodGet) && !shared.IsMethod(r, http.MethodHead) {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	video := r.PathValue("video")

	if err := store_video.Serve(video, w, r); err != nil {
		SendError(w, err)
	}
}

func AddStoreEndpoints(mux *http.ServeMux) {
	withContext := Middleware(ProxyAuthContext)
	withStore := Middleware(ProxyAuthContext, StoreContext, StoreRequired)

	mux.HandleFunc("/v0/store/user", withStore(handleStoreUser))
	mux.HandleFunc("/v0/store/magnets", withStore(handleStoreMagnets))
	mux.HandleFunc("/v0/store/magnets/check", withStore(handleStoreMagnetsCheck))
	mux.HandleFunc("/v0/store/magnets/{magnetId}", withStore(handleStoreMagnet))
	mux.HandleFunc("/v0/store/link/generate", withStore(handleStoreLinkGenerate))
	mux.HandleFunc("/v0/store/link/access/{token}", withContext(handleStoreLinkAccess))
	mux.HandleFunc("/v0/store/link/access/{token}/{filename}", withContext(handleStoreLinkAccess))

	mux.HandleFunc("/v0/store/_/static/{video}", handleStatic)
}
