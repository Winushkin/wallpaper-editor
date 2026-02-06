package main

import (
	"fmt"
	"net/http"
	"wp-editor/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/dogs", handlers.DogsHandler)

	fmt.Println("server is started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("server is dead")
	}
}
