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
	DoAuthentication = "do-authentication"
)

type DoAuthenticationRequest struct {
	PresentationRequestId int    `json:"presentationRequestId"`
	WalletId              int    `json:"walletId"`
	HolderDid             string `json:"holderDid"`
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

func (a Authnull) CallAuthService(payload *DoAuthenticationRequest) *DoAuthenticationResponse {
	fmt.Println("Calling Authnull.DoAuthentication")
	url := a.AuthNBasePath + DoAuthentication
	client := &http.Client{}
	jsonPayload, err := json.Marshal(payload)
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
