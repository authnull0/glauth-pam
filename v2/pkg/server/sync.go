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

type UserRequest struct {
	DomainId 	int    	`json:"domainId"`
	PageId   	int  	`json:"pageId"`
	PageSize 	int		`json:"pageSize"`
	Filter   	string 	`json:"filter"`
	InstanceIds	[]int  	`json:"instanceIds"`
}


var authnull0 authnull.Authnull

type Sync struct{} 
func (s *Sync) synchronize() {
	fmt.Println(" Hai Cron")
	s.pullUsers()
}

func (s *Sync) pullUsers() {

	postBody, _ := json.Marshal(&UserRequest{ 
		DomainId: 1,
		PageId: 1,
		PageSize: 10,
		Filter: "ssh",
		InstanceIds: []int{1,2},
    })
	responseBody := bytes.NewBuffer(postBody)
    resp, err := http.Post("https://api.authnull.kloudlearn.com/api/v1/lums/listAllEpmUsers", "application/json", responseBody)
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
	
	var synchronizer Sync
	authnull0.Init()
	c.AddFunc("@every 10s", func(){ synchronizer.synchronize()})
	c.Start()
}