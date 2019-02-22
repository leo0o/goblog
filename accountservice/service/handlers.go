package service

import (
	"github.com/leo0o/goblog/accountservice/dbclient"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"fmt"
)

var DBClient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter , r *http.Request) {
	var accountId = mux.Vars(r)["accountId"]

	//Read the account struct BoltDB
	account, err := DBClient.QueryAccount(accountId)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}