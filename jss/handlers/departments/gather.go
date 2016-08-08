package jss

import (
	"encoding/xml"
	"github.com/mosen/putj/jss"
	"net/http"
)

type Department struct {
	XMLName xml.Name `xml:"department" json:"-"`
	Id int `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type listResponse struct {
	XMLName xml.Name `xml:"departments"`
	Size int `xml:"size"`
	Departments []Department `xml:"department"`
}


func list(api *jss.Api) ([]Department, error) {
	req, err := api.NewRequest("GET", "/JSSResource/departments", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	deptsResponse := listResponse{}
	if err := xml.NewDecoder(res.Body).Decode(&deptsResponse); err != nil {
		return nil, err
	}

	return deptsResponse.Departments, nil
}

func DepartmentCaptureHandler(api *jss.Api, state map[string]interface{}) error {
	depts, err := list(api)
	if err != nil {
		return err
	}

	state["departments"] = depts

	return nil
}

func init() {
	jss.RegisterCaptureHandler("departments", DepartmentCaptureHandler)
}