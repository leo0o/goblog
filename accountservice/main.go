package main

import (
	"fmt"
	"github.com/leo0o/goblog/accountservice/service"
	"github.com/leo0o/goblog/accountservice/dbclient"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v...\n", appName)
	initBoltClient()
	service.StartWebServer("6767")
}


func initBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
