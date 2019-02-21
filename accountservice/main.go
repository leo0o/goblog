package main

import (
	"fmt"
	"github.com/leo0o/goblog/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v...\n", appName)
	service.StartWebServer("6767")
}
