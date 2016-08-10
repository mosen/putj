package activation

import "encoding/xml"

type ActivationCode struct {
	XMLName xml.Name `xml:"activation_code" json:"-"`
	OrganizationName string `xml:"organization_name" json:"organization_name"`
	Code string `xml:"code" json:"code"`
}
