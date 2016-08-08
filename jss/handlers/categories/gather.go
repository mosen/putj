package jss

import (
	"encoding/xml"
	"github.com/mosen/putj/jss"
	"net/http"
	"fmt"
)

type Category struct {
	XMLName xml.Name `xml:"category" json:"-"`
	Id int `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Priority int `json:"priority" xml:"priority"`
}

type listResponse struct {
	XMLName xml.Name `xml:"categories"`
	Size int `xml:"size"`
	Categories []Category `xml:"category"`
}


func detail(api *jss.Api, category *Category) error {
	req, err := api.NewRequest("GET", fmt.Sprintf("/JSSResource/categories/id/%d", category.Id), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if err := xml.NewDecoder(res.Body).Decode(category); err != nil {
		return err
	}

	return nil
}

func list(api *jss.Api) ([]Category, error) {
	req, err := api.NewRequest("GET", "/JSSResource/categories", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	categoriesResponse := listResponse{}
	if err := xml.NewDecoder(res.Body).Decode(&categoriesResponse); err != nil {
		return nil, err
	}

	return categoriesResponse.Categories, nil
}

func CategoryCaptureHandler(api *jss.Api, state map[string]interface{}) error {
	categories, err := list(api)
	if err != nil {
		return err
	}

	var catDetails []Category = make([]Category, len(categories))
	for i, cat := range categories {
		catDetail := Category{Id: cat.Id}
		if err := detail(api, &catDetail); err != nil {
			return err
		}

		catDetails[i] = catDetail
		fmt.Printf("%v\n", catDetail)
	}

	state["categories"] = categories

	return nil
}

func init() {
	jss.RegisterCaptureHandler("categories", CategoryCaptureHandler)
}