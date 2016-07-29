package jss

import "encoding/xml"

type ActivationCode struct {
	XMLName xml.Name
	ActivationCode string `xml:"activation_code"`
}
