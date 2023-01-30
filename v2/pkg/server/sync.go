package server


import (
	"gopkg.in/robfig/cron.v2"
)
func synchronize() {
	
}

func pullUsers() {}

func addToConfig() {}

func run() {

	c:= cron.New

	c.AddFunc("@every 10m", synchronize())
	c.Start()
}