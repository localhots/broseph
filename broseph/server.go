package broseph

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func newSessionHandler(rw http.ResponseWriter, req *http.Request) {
	s := NewSession()
	fmt.Fprintf(rw, s.Id)
}

func sessionListHandler(rw http.ResponseWriter, req *http.Request) {
	ids := make([]string, len(sessions))
	for i, s := range sessions {
		ids[i] = s.Id
	}
	result, _ := json.Marshal(ids)
	fmt.Fprintf(rw, string(result))
}

// Starts Broseph HTTP server on port 1879
func StartServer() {
	http.HandleFunc("/new", newSessionHandler)
	http.HandleFunc("/list", sessionListHandler)
	http.ListenAndServe(":1879", nil)
}
