package stremio_wrap

import (
	"strings"

	"github.com/SillyHippy/stremthru/store"
	"github.com/SillyHippy/stremthru/stremio"
)

func getManifest(upstream *stremio.Manifest, ud *UserData) *stremio.Manifest {
	manifest := upstream
	manifest.ID = "st:wrap:" + manifest.ID
	storeHint := ""
	if ud.StoreName == "" {
		storeHint = "ST|"
	} else {
		storeHint = strings.ToUpper(string(store.StoreName(ud.StoreName).Code())) + "|"
	}
	manifest.Name = "StremThru(" + storeHint + manifest.Name + ")"
	manifest.BehaviorHints = &stremio.BehaviorHints{
		Configurable:          true,
		ConfigurationRequired: !ud.HasRequiredValues(),
	}
	return manifest
}
