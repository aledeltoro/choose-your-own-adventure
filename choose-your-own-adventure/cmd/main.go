package main

import (
	"gophercises/choose-your-own-adventure/internal/handler"
	"gophercises/choose-your-own-adventure/internal/models"
	"gophercises/choose-your-own-adventure/internal/service"
	"gophercises/choose-your-own-adventure/internal/templating"
	"log"
	"net/http"
	"os"
)

func main() {
	data, err := os.ReadFile("assets/gopher.json")
	if err != nil {
		log.Fatalln(err)
	}

	input, err := models.NewAdventureInput(data)
	if err != nil {
		log.Fatalln(err)
	}

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	templates, err := templating.LoadTemplates()
	if err != nil {
		log.Fatalln(err)
	}

	service := service.NewStoryService(templates, input)

	handler := handler.NewHandler(service)

	http.HandleFunc("/cyoa/", handler.HandleRenderStory())

	log.Println("listening to port :8080")

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
