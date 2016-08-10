package distributionpoints

import (
	"encoding/xml"
	"net/url"
	"crypto/x509"
	"github.com/mosen/putj/jss"
	"fmt"
	"net/http"
)

type DistributionPoint struct {
	XMLName xml.Name `xml:"distribution_point" json:"-"`
	Id int `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	IPAddress string `json:"ip_address" xml:"ip_address"` // actually hostname too
	IsMaster bool `json:"is_master" xml:"is_master"`
	// FailoverPoint
	// FailoverPointURL
	EnableLoadBalancing bool `json:"enable_load_balancing,omitempty" xml:"enable_load_balancing"`
	// LocalPath
	SSHUsername string `json:"ssh_username,omitempty" xml:"ssh_username"`
	SSHPassword string `json:"ssh_password_sha256,omitempty" xml:"ssh_password_sha256"`
	ConnectionType string `json:"connection_type" xml:"connection_type"`
	ShareName string `json:"share_name" xml:"share_name"`
	WorkgroupOrDomain string `json:"workgroup_or_domain" xml:"workgroup_or_domain"`
	SharePort int `json:"share_port" xml:"share_port"`
	ReadOnlyUsername string `json:"read_only_username" xml:"read_only_username"`
	ReadOnlyPassword string `json:"read_only_password_sha256" xml:"read_only_password_sha256"`
	ReadWriteUsername string `json:"read_write_username" xml:"read_write_username"`
	ReadWritePassword string `json:"read_write_password_sha256" xml:"read_write_password_sha256"`
	HttpDownloadsEnabled bool `json:"http_downloads_enabled" xml:"http_downloads_enabled"`
	HttpURL url.URL `json:"http_url" xml:"http_url"`
	Context string
	Protocol string `json:"protocol" xml:"protocol"`
	Port int `json:"port" xml:"port"`
	NoAuthenticationRequired bool `json:"no_authentication_required" xml:"no_authentication_required"`
	UsernamePasswordRequired bool `json:"username_password_required" xml:"username_password_required"`
	HTTPUsername string `json:"http_username" xml:"http_username"`
	HTTPPassword string `json:"http_password_sha256" xml:"http_password_sha256"`
	CertificateRequired bool `json:"certificate_required" xml:"certificate_required"`
	Certificate x509.Certificate `json:"certificate, omitempty" xml:"certificate"`
}

type listResponse struct {
	XMLName xml.Name `xml:"distribution_points"`
	Size int `xml:"size"`
	DistributionPoints []DistributionPoint `xml:"distribution_point"`
}


func detail(api *jss.Api, dp *DistributionPoint) error {
	req, err := api.NewRequest("GET", fmt.Sprintf("/JSSResource/distributionpoints/id/%d", dp.Id), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if err := xml.NewDecoder(res.Body).Decode(dp); err != nil {
		return err
	}

	return nil
}

func list(api *jss.Api) ([]DistributionPoint, error) {
	req, err := api.NewRequest("GET", "/JSSResource/distributionpoints", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	dpResponse := listResponse{}
	if err := xml.NewDecoder(res.Body).Decode(&dpResponse); err != nil {
		return nil, err
	}

	return dpResponse.DistributionPoints, nil
}

func DistributionPointCaptureHandler(api *jss.Api, state map[string]interface{}) error {
	dps, err := list(api)
	if err != nil {
		return err
	}

	var dpDetails []DistributionPoint = make([]DistributionPoint, len(dps))
	for i, dp := range dps {
		dpDetail := DistributionPoint{Id: dp.Id}
		if err := detail(api, &dpDetail); err != nil {
			return err
		}

		dpDetails[i] = dpDetail
		fmt.Printf("%v\n", dpDetail)
	}

	state["distribution_points"] = dpDetails

	return nil
}

func init() {
	jss.RegisterCaptureHandler("distribution_points", DistributionPointCaptureHandler)
}