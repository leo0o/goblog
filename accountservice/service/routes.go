package service

import "net/http"

type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:"GetAccount",
		Method:"GET",
		Pattern:"/accounts/{accountId}",
		HandleFunc: func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
			writer.Write([]byte(`{"result":"OK"}`))
		},
	},
}
