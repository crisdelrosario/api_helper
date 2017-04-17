package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type API struct {
	Host        string
	ContentType string
}

type DummyPayload struct {
}

func New(host string, contentType string) API {
	return API{
		Host:        host,
		ContentType: contentType,
	}
}

func (api *API) Get(path string, auth *Auth) (*http.Response, error) {
	return api.New(http.MethodGet, path, DummyPayload{}, auth)
}

func (api *API) Post(path string, payload interface{}, auth *Auth) (*http.Response, error) {
	return api.New(http.MethodPost, path, payload, auth)
}

func (api *API) Patch(path string, payload interface{}, auth *Auth) (*http.Response, error) {
	return api.New(http.MethodPatch, path, payload, auth)
}

func (api *API) Put(path string, payload interface{}, auth *Auth) (*http.Response, error) {
	return api.New(http.MethodPut, path, payload, auth)
}

func (api *API) Delete(path string, payload interface{}, auth *Auth) (*http.Response, error) {
	return api.New(http.MethodDelete, path, payload, auth)
}

func (api *API) New(method string, path string, payload interface{}, auth *Auth) (*http.Response, error) {
	var response *http.Response
	var request *http.Request
	var err error

	data := ""
	if api.ContentType == ContentTypeJSON {
		data = api.toJSON(payload)
	}

	reader := strings.NewReader(data)

	if request, err = http.NewRequest(method, path, reader); err != nil {
		return response, err
	}

	if request == nil {
		return response, errors.New("HTTP request failed")
	}

	if auth != nil {
		request.SetBasicAuth(auth.User, auth.Pass)
	}

	request.Header.Set("Content-type", api.ContentType)
	if response, err = http.DefaultClient.Do(request); err != nil {
		return response, err
	}

	if response == nil {
		return response, errors.New("No response from the server")
	}

	return response, nil
}

func (api *API) toJSON(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
