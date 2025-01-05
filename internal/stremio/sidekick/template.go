package stremio_sidekick

import (
	"bytes"
	"embed"
	"html/template"
	"net/url"
	"strings"

	"github.com/SillyHippy/stremthru/pkg/config"
	"github.com/SillyHippy/stremthru/internal/stremio/api"
)

//go:embed *.html
var templateFs embed.FS

type TemplateData struct {
	Title          string
	Description    string
	Version        string
	IsAuthed       bool
	Email          string
	Addons         []stremio_api.Addon
	AddonOperation string
	AddonError     string
	LastAddonIndex int
	LoginMethod    string
	Login          struct {
		Email    string
		Password string
		Token    string
		Error    struct {
			Email    string
			Password string
			Token    string
		}
	}
}

type TemplateExecutor func(data *TemplateData, name string) (bytes.Buffer, error)

var funcMap = template.FuncMap{
	"url_path_escape": func(value string) string {
		return url.PathEscape(value)
	},
	"has_prefix": func(value, prefix string) bool {
		return strings.HasPrefix(value, prefix)
	},
	"get_configure_url": func(value stremio_api.Addon) string {
		if value.Manifest.BehaviorHints != nil && value.Manifest.BehaviorHints.Configurable {
			return strings.Replace(value.TransportUrl, "/manifest.json", "/configure", 1)
		}
		return ""
	},
}

var ExecuteTemplate = func() TemplateExecutor {
	tmpl := template.Must(template.New("stremio/sidekick").Funcs(funcMap).ParseFS(templateFs, "*.html"))
	return func(data *TemplateData, name string) (bytes.Buffer, error) {
		data.Version = config.Version
		if data.Addons == nil {
			data.Addons = []stremio_api.Addon{}
		}
		if data.AddonOperation == "" {
			data.AddonOperation = "move"
		}
		data.LastAddonIndex = len(data.Addons) - 1
		var buf bytes.Buffer
		err := tmpl.ExecuteTemplate(&buf, name, data)
		return buf, err
	}
}()

type PageGetter func(data *TemplateData) (bytes.Buffer, error)

var GetPage = func() PageGetter {
	return func(data *TemplateData) (bytes.Buffer, error) {
		return ExecuteTemplate(data, "index.html")
	}
}()
