package jss

import "fmt"

type Handler interface {

}

type EnforceHandlerFunc func (api *Api, state map[string]interface{}) (map[string]string, error)
var enforceHandlers map[string]EnforceHandlerFunc = make(map[string]EnforceHandlerFunc)

type CaptureHandlerFunc func (api *Api, state map[string]interface{}) error
var captureHandlers map[string]CaptureHandlerFunc = make(map[string]CaptureHandlerFunc)

func RegisterEnforceHandler(name string, handler EnforceHandlerFunc) {
	fmt.Printf("Registering enforce handler %s.\n", name)
	enforceHandlers[name] = handler
}

func RegisterCaptureHandler(name string, handler CaptureHandlerFunc) {
	fmt.Printf("Registering capture handler %s.\n", name)
	captureHandlers[name] = handler
}
