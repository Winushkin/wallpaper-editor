package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var dogApiURL string = "https://dog.ceo/api/breeds/image/random"

type apiResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func GetDogPictureURL() (string, error) {
	resp, err := http.Get(dogApiURL)
	if err != nil {
		return "", fmt.Errorf("failed to get dog API: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var ApiResp apiResponse
	if err = json.Unmarshal(body, &ApiResp); err != nil {
		return "", fmt.Errorf("cannot Unmarshall JSON Response: %w", err)
	}

	if ApiResp.Status != "success" {
		return "", fmt.Errorf("response status is not success: %w", err)
	}

	return ApiResp.Message, nil
}
