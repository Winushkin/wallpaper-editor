package main

import (
	"fmt"
	"net/http"
	"wp-editor/handlers"
)

func main() {

	http.HandleFunc("/", handlers.MainHandler)

	fmt.Println("server is started on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server is dead")
	}
}
