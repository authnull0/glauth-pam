package server


import (
	"gopkg.in/robfig/cron.v2"
	"fmt"
)

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

	
	c.AddFunc("@every 1s", func(){ synchronize()})
	c.Start()
}