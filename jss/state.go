package jss

// Unmarshal everything as a byte slice so that each handler can decode its own struct
type State struct {
	ActivationCode []byte `json:"activation_code,omitempty"`
}
