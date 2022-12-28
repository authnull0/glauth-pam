package external

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type OtpResposne struct {
	Status string `json:"status"`
	code   int    `json:"code"`
}

func OtpCaller(username string, phoneNo string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"username":    "Hussain",
		"entity":      "+919962616643",
		"entity_type": "otp",
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("http://localhost:2882/authnull0/api/v1/authn/authenticate-user", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	var res OtpResposne
	json.NewDecoder(resp.Body).Decode(&res)

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
	return true
}
