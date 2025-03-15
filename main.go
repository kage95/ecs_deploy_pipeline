package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api/hello", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := Response{
		Message: "Hello",
	}
	json.NewEncoder(w).Encode(res)
}
