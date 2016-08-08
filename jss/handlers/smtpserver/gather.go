package smtpserver

import (
	"encoding/xml"
	"github.com/mosen/putj/jss"
	"net/http"
	"fmt"
)

type SmtpServer struct {
	XMLName xml.Name `xml:"smtp_server" json:"-"`
	Enabled bool `xml:"enabled" json:"enabled"`
	Host string `xml:"host" json:"host"`
	Port int `xml:"port" json:"port"`
	Timeout int `xml:"timeout" json:"timeout"`
	AuthorizationRequired bool `xml:"authorization_required" json:"authorization_required"`
	Username string `xml:"username,omitempty" json:"username,omitempty"`
	PasswordHash string `xml:"password_sha256" json:"password_sha256"`
	SSL bool `xml:"ssl" json:"ssl"`
	TLS bool `xml:"tls" json:"tls"`
	SendFromName string `xml:"send_from_name" json:"send_from_name"`
	SendFromEmail string `xml:"send_from_email" json:"send_from_email"`
}

func SmtpServerCaptureHandler(api *jss.Api, state map[string]interface{}) error {
	req, err := api.NewRequest("GET", "/JSSResource/smtpserver", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	smtpServer := &SmtpServer{}
	if err := xml.NewDecoder(res.Body).Decode(smtpServer); err != nil {
		return err
	}

	state["smtp_server"] = smtpServer
	fmt.Printf("%v\n", smtpServer)

	return nil
}

func init() {
	jss.RegisterCaptureHandler("smtp_server", SmtpServerCaptureHandler)
}



