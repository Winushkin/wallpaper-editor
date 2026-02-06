package handlers

import (
	"context"
	"go.uber.org/zap"
	"html/template"
	"io"
	"net/http"
	"os"
	"wp-editor/internal/services"
	"wp-editor/pkg/logger"
)

type DogPageData struct {
	ImageURL string
}

var mainTemplate = template.Must(template.ParseFiles("templates/index.html"))
var dogsTemplate = template.Must(template.ParseFiles("templates/dogs.html"))

func MainHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = mainTemplate.Execute(w, map[string]any{
		"Title": "Собачки!!!",
	})
}

func DogsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = logger.GetContextWithNewLogger(ctx, true)
	log := logger.GetLoggerFromCtx(ctx)

	log.Debug(ctx, "Начало запроса")

	imageURL, err := services.GetDogPictureURL()
	if err != nil {
		log.Error(ctx, "GetDogPictureURL:", zap.Error(err))
		return
	}

	err = dogsTemplate.Execute(w, DogPageData{
		ImageURL: imageURL,
	})
	if err != nil {
		return
	}
	log.Debug(ctx, "Конец запроса")
}

func serverFallback(w http.ResponseWriter) {
	f, _ := os.Open("dropPhoto.jpg")
	defer f.Close()

	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, f)
}
