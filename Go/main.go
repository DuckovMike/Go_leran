package main

import (
	"encoding/json"
	"net/http"
)

type HttpResp struct {
	Response string
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	resp, _ := json.Marshal(HttpResp{Response: "hello"})
	w.Write(resp)
}

func main() {
	http.HandleFunc("/hello", helloHandle)

	http.ListenAndServe(":8080", nil)
}
