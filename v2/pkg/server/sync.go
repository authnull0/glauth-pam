package server


import (
	"gopkg.in/robfig/cron.v2"
	"fmt"
	authnull "github.com/glauth/glauth/v2/pkg/external"
)

var authnull0 authnull.Authnull

type Sync struct{} 
func synchronize() {
	fmt.Println(" Hai Cron")
}

func pullUsers() {

}

func addToConfig() {

}

func (s Sync)Run() {

	c:= cron.New()
	
	
	authnull0.Init()
	c.AddFunc("@every 1s", func(){ synchronize()})
	c.Start()
}