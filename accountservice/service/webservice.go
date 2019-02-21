package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	log.Println("starting http service at port " + port)
	r := NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
