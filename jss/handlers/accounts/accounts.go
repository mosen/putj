package jss

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"github.com/mosen/putj/jss"
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
	XMLName xml.Name `json:"-"`
	Id int `xml:"Id" json:"id"`
	Name string `xml:"name" json:"name,omitempty"`
	DirectoryUser bool `xml:"directory_user" json:"directory_user,omitempty"`
	FullName string `xml:"full_name" json:"full_name,omitempty"`
	Email string `xml:"email" json:"email,omitempty"`
	Password string `xml:"password_sha256" json:"password_sha256,omitempty"`
	//PasswordSince string `xml:"password_sha256`
	AccessLevel AccessLevel `xml:"access_level" json:"access_level"`
	PrivilegeSet PrivilegeSet `xml:"privilege_set" json:"privilege_set"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Id int `xml:"id"`
	Name string `xml:"name,omitempty"`
}

type Group struct {
  	XMLName xml.Name `xml:"group"`
	Id int `xml:"id"`
	Name string `xml:"name"`
	SiteId int `xml:"site>id"`
	SiteName string `xml:"site>name"`
}

type AccountsList struct {
	XMLName xml.Name `xml:"accounts" json:"-"`
	Users []User `xml:"users>user" json:"users,omitempty"`
	Groups []Group `xml:"groups>group,omitempty" json:"groups,omitempty"`
}

type Result struct {
	Users []Account `json:"accounts"`
}

func list(api *jss.Api) (*AccountsList, error) {
	req, err := api.NewRequest("GET", "/JSSResource/accounts", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	accounts := AccountsList{}
	if err := xml.NewDecoder(res.Body).Decode(&accounts); err != nil {
		return nil, err
	}

	return &accounts, nil
}

func detail(api *jss.Api, account *Account) error {
	req, err := api.NewRequest("GET", fmt.Sprintf("/JSSResource/accounts/userid/%d", account.Id), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if err := xml.NewDecoder(res.Body).Decode(account); err != nil {
		return err
	}

	return nil
}


func AccountsCaptureHandler(api *jss.Api, state map[string]interface{}) error {
	accounts, err := list(api)
	if err != nil {
		return err
	}

	var userDetails []Account = make([]Account, len(accounts.Users))
	for i, user := range accounts.Users {
		account := Account{Id: user.Id}
		if err := detail(api, &account); err != nil {
			return err
		}

		userDetails[i] = account
		fmt.Printf("%v\n", account)
	}

	state["users"] = userDetails

	return nil
}


func init() {
	//RegisterEnforceHandler("accounts", AccountsEnforceHandler)
	jss.RegisterCaptureHandler("accounts", AccountsCaptureHandler)
}