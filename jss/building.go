package jss

import "encoding/xml"

type Building struct {
	XMLName xml.Name
	Name string `json:"name" xml:"name"`
}
