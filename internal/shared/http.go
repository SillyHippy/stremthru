package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/SillyHippy/stremthru/core"
	"github.com/SillyHippy/stremthru/internal/request"
)

func GetHTTPClient(withProxy bool) *http.Client {
	transport := request.DefaultHTTPTransport.Clone()
	if !withProxy {
		transport.Proxy = nil
	}
	return &http.Client{
		Transport: transport,
		Timeout:   90 * time.Second,
	}
}

func IsMethod(r *http.Request, method string) bool {
	return r.Method == method
}

func GetQueryInt(queryParams url.Values, name string, defaultValue int) (int, error) {
	if qVal, ok := queryParams[name]; ok {
		v := qVal[0]
		if v == "" {
			return defaultValue, nil
		}

		val, err := strconv.Atoi(v)
		if err != nil {
			return 0, errors.New("invalid " + name)
		}
		return val, nil
	}
	return defaultValue, nil
}

func ReadRequestBodyJSON[T interface{}](r *http.Request, payload T) error {
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return ErrorUnsupportedMediaType(r)
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err == nil {
		return err
	}
	if err == io.EOF {
		return ErrorBadRequest(r, "missing body")
	}
	error := core.NewAPIError("failed to decode body")
	error.Cause = err
	return error
}

type response struct {
	Data  any         `json:"data,omitempty"`
	Error *core.Error `json:"error,omitempty"`
}

func (r response) send(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Printf("failed to encode json %v\n", err)
	}
}

func SendError(w http.ResponseWriter, err error) {
	isUnexpected := false

	var e core.StremThruError
	if sterr, ok := err.(core.StremThruError); ok {
		e = sterr
	} else {
		e = &core.Error{Cause: err}
		isUnexpected = true
	}
	e.Pack()

	if isUnexpected {
		log.Printf("unexpected error: %v\n", err)
	}

	res := &response{}
	res.Error = e.GetError()

	res.send(w, e.GetStatusCode())
}

func SendResponse(w http.ResponseWriter, statusCode int, data any, err error) {
	if err != nil {
		SendError(w, err)
		return
	}

	res := &response{}
	res.Data = data

	res.send(w, statusCode)
}

func SendHTML(w http.ResponseWriter, statusCode int, data bytes.Buffer) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)
	data.WriteTo(w)
}

func copyHeaders(src http.Header, dest http.Header) {
	for key, values := range src {
		for _, value := range values {
			dest.Add(key, value)
		}
	}
}

var proxyHttpClientWithTunnel = func() *http.Client {
	transport := request.DefaultHTTPTransport.Clone()
	return &http.Client{
		Transport: transport,
	}
}()

var proxyHttpClientWithoutTunnel = func() *http.Client {
	transport := request.DefaultHTTPTransport.Clone()
	transport.Proxy = nil
	return &http.Client{
		Transport: transport,
	}
}()

func ProxyResponse(w http.ResponseWriter, r *http.Request, url string, useTunnel bool) {
	request, err := http.NewRequest(r.Method, url, nil)
	if err != nil {
		e := ErrorInternalServerError(r, "failed to create request")
		e.Cause = err
		SendError(w, e)
		return
	}

	copyHeaders(r.Header, request.Header)

	proxyHttpClient := proxyHttpClientWithTunnel
	if !useTunnel {
		proxyHttpClient = proxyHttpClientWithoutTunnel
	}

	response, err := proxyHttpClient.Do(request)
	if err != nil {
		e := ErrorBadGateway(r, "failed to request url")
		e.Cause = err
		SendError(w, e)
		return
	}
	defer response.Body.Close()

	copyHeaders(response.Header, w.Header())

	w.WriteHeader(response.StatusCode)

	_, err = io.Copy(w, response.Body)
	if err != nil {
		log.Printf("stream failure: %v", err)
	}
}

func extractRequestScheme(r *http.Request) string {
	scheme := r.Header.Get("X-Forwarded-Proto")

	if scheme == "" {
		scheme = r.URL.Scheme
	}

	if scheme == "" {
		scheme = "http"
		if r.TLS != nil {
			scheme = "https"
		}
	}

	return scheme
}

func extractRequestHost(r *http.Request) string {
	host := r.Header.Get("X-Forwarded-Host")

	if host == "" {
		host = r.Host
	}

	return host
}

func ExtractRequestBaseURL(r *http.Request) *url.URL {
	return &url.URL{
		Scheme: extractRequestScheme(r),
		Host:   extractRequestHost(r),
	}
}
