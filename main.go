package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type randomResp struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("Начало запроса")
		resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var respStruct randomResp
		if err = json.Unmarshal(body, &respStruct); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if respStruct.Status != "success" {
			http.Error(w, "dog.ceo api error", 500)
			return
		}

		imgResp, err := http.Get(respStruct.Message)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		defer imgResp.Body.Close()

		w.Header().Set("Content-Type", imgResp.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusOK)

		_, err = io.Copy(w, imgResp.Body)
		if err != nil {
			fmt.Printf("streaming error: %v", err)
		}
		fmt.Println("Конец запроса")
	})

	fmt.Println("server is started on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server is dead")
	}
}
