package endpoint

import (
	"net/http"
	"net/url"

	"github.com/SillyHippy/stremthru/internal/shared"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		shared.ErrorMethodNotAllowed(r).Send(w)
		return
	}

	targetUrl := r.URL.Query().Get("url")
	if targetUrl == "" {
		shared.ErrorBadRequest(r, "missing url").Send(w)
		return
	}

	targetUrl, err := url.QueryUnescape(targetUrl)
	if err != nil {
		e := shared.ErrorBadRequest(r, "invalid url")
		e.Cause = err
		e.Send(w)
		return
	}

	if u, err := url.ParseRequestURI(targetUrl); err != nil || u.Scheme == "" || u.Host == "" {
		e := shared.ErrorBadRequest(r, "invalid url")
		e.Cause = err
		e.Send(w)
		return
	}

	shared.ProxyResponse(w, r, targetUrl, true)
}

func AddProxyEndpoints(mux *http.ServeMux) {
	withMiddleware := Middleware(ProxyAuthContext, ProxyAuthRequired)

	mux.HandleFunc("/proxy", withMiddleware(handleProxy))
	mux.HandleFunc("/v0/proxy", withMiddleware(handleProxy))
}
