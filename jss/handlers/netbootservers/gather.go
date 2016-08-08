package netbootservers

import (
	"encoding/xml"
	"fmt"
	"github.com/mosen/putj/jss"
	"net/http"
)

type NetbootServersList struct {
	XMLName xml.Name `xml:"netboot_servers" json:"-"`
	Servers []NetbootServerSummary `xml:"netboot_server"`
}

type NetbootServerSummary struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

type NetbootServer struct {
	XMLName           xml.Name   `xml:"netboot_server" json:"-"`
	Id                int        `xml:"id" json:"id"`
	Name              string     `xml:"name" json:"name"`
	IpAddress         string `xml:"ip_address" json:"ip_address"`
	DefaultImage      bool       `xml:"default_image" json:"default_image"`
	SpecificImage     bool       `xml:"specific_image" json:"specific_image"`
	TargetPlatform    string     `xml:"target_platform" json:"target_platform"`
	SharePoint        string     `xml:"share_point" json:"share_point"`
	Set               string     `xml:"set" json:"set"`
	Image             string     `xml:"image" json:"image"`
	Protocol          string     `xml:"protocol" json:"protocol"`
	ConfigureManually bool       `xml:"configure_manually" json:"configure_manually"`
	BootArgs          string     `xml:"boot_args" json:"boot_args"`
	BootFile          string     `xml:"boot_file" json:"boot_file"`
	BootDevice        string     `xml:"boot_device" json:"boot_device"`
}

type Result struct {
	NetbootServers []NetbootServer `json:"netboot_servers"`
}

func list(api *jss.Api) (*NetbootServersList, error) {
	req, err := api.NewRequest("GET", "/JSSResource/netbootservers", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	servers := NetbootServersList{}
	if err := xml.NewDecoder(res.Body).Decode(&servers); err != nil {
		return nil, err
	}

	return &servers, nil
}

func detail(api *jss.Api, server *NetbootServer) error {
	req, err := api.NewRequest("GET", fmt.Sprintf("/JSSResource/netbootservers/id/%d", server.Id), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if err := xml.NewDecoder(res.Body).Decode(server); err != nil {
		return err
	}

	return nil
}


func NetbootServersCaptureHandler(api *jss.Api, state map[string]interface{}) error {
	servers, err := list(api)
	if err != nil {
		return err
	}

	var serverDetails []NetbootServer = make([]NetbootServer, len(servers.Servers))
	for i, server := range servers.Servers {
		serverDetail := NetbootServer{Id: server.Id}
		if err := detail(api, &serverDetail); err != nil {
			return err
		}

		serverDetails[i] = serverDetail
		fmt.Printf("%v\n", serverDetail)
	}

	state["netboot_servers"] = serverDetails

	return nil
}

func init() {
	jss.RegisterCaptureHandler("netbootservers", NetbootServersCaptureHandler)
}
