package main

import (
	"fmt"
	"net/http"
)

type randomResp struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {

	http.HandleFunc("/", MainHandler)

	fmt.Println("server is started on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server is dead")
	}
}
