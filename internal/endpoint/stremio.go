package endpoint

import (
	"net/http"

	"github.com/SillyHippy/stremthru/internal/config"
	"github.com/SillyHippy/stremthru/internal/stremio/disabled"
	"github.com/SillyHippy/stremthru/internal/stremio/sidekick"
	"github.com/SillyHippy/stremthru/internal/stremio/store"
	"github.com/SillyHippy/stremthru/internal/stremio/wrap"
)

func AddStremioEndpoints(mux *http.ServeMux) {
	if config.StremioAddon.IsEnabled("store") {
		stremio_store.AddStremioStoreEndpoints(mux)
	}
	if config.StremioAddon.IsEnabled("wrap") {
		stremio_wrap.AddStremioWrapEndpoints(mux)
	}
	if config.StremioAddon.IsEnabled("sidekick") {
		stremio_sidekick.AddStremioSidekickEndpoints(mux)
		stremio_disabled.AddStremioDisabledEndpoints(mux)
	}
}
