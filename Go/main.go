package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpResponse struct {
	Fisrt  string
	Second string
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	fisrt := r.URL.Query().Get("f")
	second := r.URL.Query().Get("s")

	resp := httpResponse{
		Fisrt:  fisrt,
		Second: second,
	}

	fmt.Println(resp)
	httpResp, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("ошибка", err)
	}

	w.Write(httpResp)

}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.ListenAndServe(":8080", nil)
}
