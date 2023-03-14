package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Authnull struct {
	AuthNBasePath string `default:"https://api.did.kloudlearn.com/authnull0/api/v1/authn" envconfig:"AUTHN_BASE_PATH"`
}

const (
	DoAuthentication = "/do-authentication"
)

type DoAuthenticationRequest struct {
	Username     string `json:"username"`
	UserSource   string `json:"userSource"`
	Endpoint     string `json:"endpoint"`
	Group        string `json:"group"`
	ResponseType string `json:"responseType"`
}

type DoAuthenticationResponse struct {
	IsValid    bool        `json:"isValid"`
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	Code       int         `json:"code"`
	Credential interface{} `json:"credential"`
}

func (a Authnull) Init() {
	fmt.Println("Init Authnull")
}

func (a Authnull) FetchUsers() {}

func (a Authnull) CallAuthService(username, groupName string) *DoAuthenticationResponse {
	fmt.Println("Calling Authnull.DoAuthentication")
	url := a.AuthNBasePath + DoAuthentication
	client := &http.Client{}
	jsonPayload, err := json.Marshal(DoAuthenticationRequest{
		Username:     username,
		UserSource:   "AD",
		Endpoint:     "172.83.61.9",
		Group:        groupName,
		ResponseType: "password",
	})
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	var res DoAuthenticationResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
	}

	return &res
}
