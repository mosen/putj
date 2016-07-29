package jss

import (
	"encoding/xml"
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
	"net/http"
	//"log"
)

type AccessLevel string
const (
	ALFull AccessLevel = "Full Access"
	ALSite = "Site Access"
)

type PrivilegeSet string
const (
	PSAdministrator PrivilegeSet = "Administrator"
	PSAuditor = "Auditor"
	PSEnroll = "Enrollment Only"
	PSCustom = "Custom"
)

// This structure describes the format of a JSS user account
type Account struct {
	XMLName xml.Name
	Name string `xml:"name"`
	DirectoryUser bool `xml:"directory_user"`
	FullName string `xml:"full_name"`
	Email string `xml:"email"`
	//Password string `xml:"password_sha256"`
	AccessLevel AccessLevel `xml:"access_level"`
	PrivilegeSet PrivilegeSet `xml:"privilege_set"`
}

//func getJssAccount(account Account) (*Account, error) {
//
//}

func postJssAccount(account Account) error {
	accountXml, err := xml.Marshal(account)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(accountXml))

	api, err := NewApi("http://localhost:9080", "admin", "password")
	body := ioutil.NopCloser(bytes.NewBuffer(accountXml))
	req, err := api.NewRequest("POST", "/JSSResource/accounts/userid/4", body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp)

	return nil
}

func AccountsEnforceHandler(api *Api, state State) (map[string]string, error) {


	if state.Accounts == nil || len(state.Accounts) == 0 {
		fmt.Println("No account(s) defined")
		return nil, nil
	}

	for _, _ = range state.Accounts {

	}

       	return nil, nil
}

func AccountsCaptureHandler(api *Api, state *State) error {
	req, err := api.NewRequest("GET", "/JSSResource/accounts", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}


func init() {
	RegisterEnforceHandler("accounts", AccountsEnforceHandler)
	RegisterCaptureHandler("accounts", AccountsCaptureHandler)
}