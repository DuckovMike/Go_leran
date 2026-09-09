package main

import (
	httpT "Go/http"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", httpT.HandleClick)
	http.ListenAndServe(":8080", nil)

	fmt.Println("200")
}
