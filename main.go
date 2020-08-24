package main

import (
	"encoding/json"
	"net/http"
)

type JsonMessage struct {
	M string `json:"message"`
}

// Ok Handler
func Ok(w http.ResponseWriter, req *http.Request) {
	ok := JsonMessage{M: "Ok"}

	j, _ := json.Marshal(ok)

	w.Header().Add("Content-Type", "application/json")

	w.Write(j)
}

func main() {
	http.HandleFunc("/ok", Ok)
	http.ListenAndServe(":8080", nil)
}
