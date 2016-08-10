package activation

import (
	"github.com/mosen/putj/jss"
	"fmt"
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"bytes"
)

func ActivationEnforceHandler(api *jss.Api, state map[string]interface{}) (map[string]string, error) {
	fmt.Printf("%v\n", state)
	s, ok := state["activation_code"]
	if (!ok) {
		return nil, nil
	}

	si := s.(map[string]interface{})

	activationCode := &ActivationCode{
		OrganizationName: si["organization_name"].(string),
		Code: si["code"].(string),
	}

	reqBody, err := xml.Marshal(&activationCode)
	if err != nil {
		return nil, err
	}

	reqBodyReader := ioutil.NopCloser(bytes.NewReader(reqBody))
	req, err := api.NewRequest("PUT", "/JSSResource/activationcode", reqBodyReader)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(res.StatusCode)

	return nil, nil
}

func init() {
	jss.RegisterEnforceHandler("activation_code", ActivationEnforceHandler)
}

