package jss

import (
	"encoding/xml"
	"github.com/mosen/putj/jss"
	"net/http"
)

type Building struct {
	XMLName xml.Name `xml:"building" json:"-"`
	Id int `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type listResponse struct {
	XMLName xml.Name `xml:"buildings"`
	Size int `xml:"size"`
	Buildings []Building `xml:"building"`
}


func list(api *jss.Api) ([]Building, error) {
	req, err := api.NewRequest("GET", "/JSSResource/buildings", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	buildingsResponse := listResponse{}
	if err := xml.NewDecoder(res.Body).Decode(&buildingsResponse); err != nil {
		return nil, err
	}

	return buildingsResponse.Buildings, nil
}

func BuildingCaptureHandler(api *jss.Api, state map[string]interface{}) error {
	buildings, err := list(api)
	if err != nil {
		return err
	}

	state["buildings"] = buildings

	return nil
}

func init() {
	jss.RegisterCaptureHandler("buildings", BuildingCaptureHandler)
}