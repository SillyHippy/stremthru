package stremio_store

import (
	"strings"

	"github.com/SillyHippy/stremthru/core"
	"github.com/SillyHippy/stremthru/internal/config"
	"github.com/SillyHippy/stremthru/store"
	"github.com/SillyHippy/stremthru/stremio"
)

const CATALOG_ID = "st:store"
const ID_PREFIX = CATALOG_ID + ":"
const STORE_ACTION_ID = ID_PREFIX + "action"
const STORE_ACTION_ID_PREFIX = STORE_ACTION_ID + ":"

const ContentTypeOther = "other"

const (
	CatalogTypeVideo  = "video"
	CatalogTypeAction = "action"
)

func getManifest(ud *UserData) *stremio.Manifest {
	name := "Store"
	description := "StremThru Store Catalog and Search"
	switch ud.StoreName {
	case "":
	case "stremthru":
	default:
		name = name + " | " + strings.ToUpper(string(store.StoreName(ud.StoreName).Code()))
		description = description + " - " + ud.StoreName
	}

	contactEmail, _ := core.Base64Decode("Z2l0aHViQG11bmlmdGFuamltLmRldg==")
	manifest := &stremio.Manifest{
		ID:          "dev.muniftanjim.stremthru.store",
		Name:        name,
		Description: description,
		Version:     config.Version,
		Resources: []stremio.Resource{
			stremio.Resource{
				Name:       stremio.ResourceNameMeta,
				Types:      []stremio.ContentType{ContentTypeOther},
				IDPrefixes: []string{ID_PREFIX},
			}, stremio.Resource{
				Name:       stremio.ResourceNameStream,
				Types:      []stremio.ContentType{ContentTypeOther},
				IDPrefixes: []string{ID_PREFIX},
			},
		},
		Types: []stremio.ContentType{},
		Catalogs: []stremio.Catalog{
			stremio.Catalog{
				Id:   CATALOG_ID,
				Name: "Store",
				Type: ContentTypeOther,
				Extra: []stremio.CatalogExtra{
					stremio.CatalogExtra{
						Name: "search",
					},
					stremio.CatalogExtra{
						Name: "skip",
					},
					stremio.CatalogExtra{
						Name:    "type",
						Options: []string{CatalogTypeVideo, CatalogTypeAction},
					},
				},
			},
		},
		ContactEmail: contactEmail,
		BehaviorHints: &stremio.BehaviorHints{
			Configurable:          true,
			ConfigurationRequired: !ud.HasRequiredValues(),
		},
	}

	return manifest
}
