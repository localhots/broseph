package broseph

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewSessionHandler(rw http.ResponseWriter, req *http.Request) {
	s := NewSession()
	fmt.Fprintf(rw, s.Id)
}

func SessionListHandler(rw http.ResponseWriter, req *http.Request) {
	ids := make([]string, len(sessions))
	for i, s := range sessions {
		ids[i] = s.Id
	}
	result, _ := json.Marshal(ids)
	fmt.Fprintf(rw, string(result))
}

func StartServer() {
	http.HandleFunc("/new", NewSessionHandler)
	http.HandleFunc("/list", SessionListHandler)
	http.ListenAndServe(":1879", nil)
}
