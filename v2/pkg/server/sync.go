package server


import (
	"bytes"
	"gopkg.in/robfig/cron.v2"
	"fmt"
   "net/http"
   "encoding/json"
	"log"
	"io/ioutil"
	authnull "github.com/glauth/glauth/v2/pkg/external"
)

var authnull0 authnull.Authnull

type Sync struct{} 
func (s *Sync) synchronize() {
	fmt.Println(" Hai Cron")
}

func (s *Sync) pullUsers() {

	postBody, _ := json.Marshal(map[string]string{
      "api-key":  "abc",
    })
	responseBody := bytes.NewBuffer(postBody)
    resp, err := http.Post("https://client.did.kloudlearn.com/api/v1/lums/listAllEpmUsers", "application/json", responseBody)
	if err != nil {
      log.Fatalf("An Error Occured %v", err)
   }
   defer resp.Body.Close()
//Read the response body
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
   fmt.Println(string(body))

}

func (s *Sync) addToConfig() {

}

func (s Sync) Run() {

	c:= cron.New()
	
	var sync Sync
	authnull0.Init()
	c.AddFunc("@every 1s", func(){ sync.synchronize()})
	c.Start()
}