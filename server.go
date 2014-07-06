package broseph

import (
	"fmt"
	"net/http"
)

func NewSessionHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "OK")
}

func SessionListHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "OK")
}

func StartServer() {
	http.HandleFunc("/new", NewSessionHandler)
	http.HandleFunc("/list", SessionListHandler)
	http.ListenAndServe(":1879", nil)
}
