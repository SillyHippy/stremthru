package endpoint

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/SillyHippy/stremthru/pkg/config"
	"github.com/SillyHippy/stremthru/internal/context"
	"github.com/SillyHippy/stremthru/internal/request"
	"github.com/SillyHippy/stremthru/internal/shared"
)

type HealthData struct {
	Status string `json:"status"`
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	health := &HealthData{}
	health.Status = "ok"
	SendResponse(w, 200, health, nil)
}

func getIp(client *http.Client) (string, error) {
	ip_req, err := http.NewRequest(http.MethodGet, "https://checkip.amazonaws.com", nil)
	if err != nil {
		return "", err
	}

	ip_res, err := client.Do(ip_req)
	if err != nil {
		return "", err
	}

	ip_body, err := io.ReadAll(ip_res.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(ip_body)), nil
}

func getMachineIp() (string, error) {
	client := shared.GetHTTPClient(false)
	return getIp(client)
}

type HealthDebugDataIP struct {
	Exposed string `json:"exposed"`
	Machine string `json:"machine"`
}

type HealthDebugDataStore struct {
	Default string   `json:"default"`
	Names   []string `json:"names"`
}

type HealthDebugDataUser struct {
	Name  string               `json:"name"`
	Store HealthDebugDataStore `json:"store"`
}

type HealthDebugData struct {
	Time    string               `json:"time"`
	Version string               `json:"version"`
	User    *HealthDebugDataUser `json:"user,omitempty"`
	IP      *HealthDebugDataIP   `json:"ip,omitempty"`
}

func handleHealthDebug(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetRequestContext(r)

	data := &HealthDebugData{
		Time:    time.Now().Format(time.RFC3339),
		Version: config.Version,
	}

	if ctx.IsProxyAuthorized {
		data.User = &HealthDebugDataUser{
			Name: ctx.ProxyAuthUser,
			Store: HealthDebugDataStore{
				Default: config.StoreAuthToken.GetPreferredStore(ctx.ProxyAuthUser),
				Names:   config.StoreAuthToken.ListStores(ctx.ProxyAuthUser),
			},
		}
		exposedIp, _ := getIp(request.DefaultHTTPClient)
		machineIp, _ := getMachineIp()
		data.IP = &HealthDebugDataIP{
			Exposed: exposedIp,
			Machine: machineIp,
		}
	}

	SendResponse(w, 200, data, nil)
}

func AddHealthEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/v0/health", handleHealth)
	mux.HandleFunc("/v0/health/__debug__", Middleware(ProxyAuthContext, StoreContext)(handleHealthDebug))
}
