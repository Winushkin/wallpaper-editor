package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type randomResp struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Начало запроса")
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		serverFallback(w)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var respStruct randomResp
	if err = json.Unmarshal(body, &respStruct); err != nil {
		serverFallback(w)
		return
	}

	if respStruct.Status != "success" {
		serverFallback(w)
		return
	}

	imgResp, err := http.Get(respStruct.Message)
	if err != nil {
		serverFallback(w)
	}
	defer imgResp.Body.Close()

	fmt.Println(imgResp.Header.Get("Content-Type"))
	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, imgResp.Body)
	if err != nil {
		serverFallback(w)
	}
	fmt.Println("Конец запроса")
}

func serverFallback(w http.ResponseWriter) {
	f, _ := os.Open("dropPhoto.jpg")
	defer f.Close()

	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, f)
}
