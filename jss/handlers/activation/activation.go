package jss

import (
	"encoding/xml"
	"github.com/mosen/putj/jss"
	"net/http"
	"fmt"
)

type ActivationCode struct {
	XMLName xml.Name `xml:"activation_code" json:"-"`
	OrganizationName string `xml:"organization_name" json:"organization_name"`
	Code string `xml:"code" json:"code"`
}

func ActivationCaptureHandler(api *jss.Api, state map[string]interface{}) error {
	req, err := api.NewRequest("GET", "/JSSResource/activationcode", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	activationCode := &ActivationCode{}
	if err := xml.NewDecoder(res.Body).Decode(activationCode); err != nil {
		return err
	}

	state["activation_code"] = activationCode
	fmt.Printf("%v\n", activationCode)

	return nil
}

func init() {
	jss.RegisterCaptureHandler("activation_code", ActivationCaptureHandler)
}


