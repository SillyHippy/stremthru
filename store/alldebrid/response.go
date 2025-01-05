package alldebrid

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/SillyHippy/stremthru/core"
)

type ResponseError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e *ResponseError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

type Response[T interface{}] struct {
	Status string         `json:"status"`
	Data   T              `json:"data,omitempty"`
	Error  *ResponseError `json:"error,omitempty"`
}

type ResponseEnvelop interface {
	GetStatus() string
	GetError() *ResponseError
}

func (r Response[any]) GetStatus() string {
	return r.Status
}

func (r Response[any]) GetError() *ResponseError {
	return r.Error
}

type APIResponse[T interface{}] struct {
	Header     http.Header
	StatusCode int
	Data       T
}

func newAPIResponse[T interface{}](res *http.Response, data T) APIResponse[T] {
	apiResponse := APIResponse[T]{
		StatusCode: 503,
		Data:       data,
	}
	if res != nil {
		apiResponse.Header = res.Header
		apiResponse.StatusCode = res.StatusCode
	}
	return apiResponse
}

func extractResponseError(statusCode int, body []byte, v ResponseEnvelop) error {
	if v.GetStatus() == "error" {
		return v.GetError()
	}
	if statusCode >= http.StatusBadRequest {
		return errors.New(string(body))
	}
	return nil
}

func processResponseBody(res *http.Response, err error, v ResponseEnvelop) error {
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	err = core.UnmarshalJSON(res.StatusCode, body, v)
	if err != nil {
		return err
	}

	return extractResponseError(res.StatusCode, body, v)
}
