package jss

import (
	"fmt"
	"net/http"
	"io"
	"net/url"
)

type Api struct {
	url      *url.URL
	username string
	password string
}

type Endpoint interface {
	Path(...string) string

	Index() []interface{}
	Get(id int) interface{}
	Put(id int) interface{}
	Post(id int) interface{}
}

func NewApi(urlString string, username string, password string) (*Api, error) {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	return &Api{
		url,
		username,
		password,
	}, nil
}

func (a *Api) NewRequest(method string, relPath string, body io.ReadCloser) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", a.url, relPath), body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(a.username, a.password)
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("Accept", "text/xml")

	return req, nil
}

func (a *Api) Enforce(state State) error {
	for name, handlerFunc := range enforceHandlers {
		fmt.Printf("Executing %s\n", name)
		if _, err := handlerFunc(a, state); err != nil {
			return err
		}
	}

	return nil
}

func (a *Api) Capture() (*State, error) {
	state := &State{}

	for name, handlerFunc := range captureHandlers {
		fmt.Printf("Executing %s\n", name)
		if err := handlerFunc(a, state); err != nil {
			return state, err
		}
	}

	return state, nil
}


type EnforceHandlerFunc func (api *Api, state State) (map[string]string, error)
var enforceHandlers map[string]EnforceHandlerFunc = map[string]EnforceHandlerFunc{}

type CaptureHandlerFunc func (api *Api, state *State) error
var captureHandlers map[string]CaptureHandlerFunc = map[string]CaptureHandlerFunc{}

func RegisterEnforceHandler(name string, handler EnforceHandlerFunc) {
	fmt.Printf("Registering enforce handler %s.\n", name)
	enforceHandlers[name] = handler
}

func RegisterCaptureHandler(name string, handler CaptureHandlerFunc) {
	fmt.Printf("Registering capture handler %s.\n", name)
	captureHandlers[name] = handler
}
