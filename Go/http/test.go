package httpT

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonR struct {
	Key string `json:"key"`
}

func HandleClick(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	rTest := JsonR{}

	err := json.NewDecoder(r.Body).Decode(&rTest)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rTest)

	_, err = w.Write([]byte(rTest.Key))
	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()
}
