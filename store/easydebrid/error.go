package easydebrid

import (
	"github.com/SillyHippy/stremthru/core"
	"github.com/SillyHippy/stremthru/store"
)

func UpstreamErrorWithCause(cause error) *core.UpstreamError {
	err := core.NewUpstreamError("")
	err.StoreName = string(store.StoreNameEasyDebrid)

	if rerr, ok := cause.(*ResponseContainer); ok {
		err.Msg = rerr.Err
		err.UpstreamCause = rerr
	} else {
		err.Cause = cause
	}

	return err
}
