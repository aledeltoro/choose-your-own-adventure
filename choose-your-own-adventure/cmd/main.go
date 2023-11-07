package main

import (
	"fmt"
	"gophercises/choose-your-own-adventure/internal/handler"
	"gophercises/choose-your-own-adventure/internal/models"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// Change the current working directory to be consistent, regardless of where the program is run
	data, err := os.ReadFile("../gopher.json")
	if err != nil {
		log.Fatalln(err)
	}

	input, err := models.NewAdventureInput(data)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl := template.Must(template.ParseFiles("../assets/html/base.html"))

	handler := handler.NewHandler(tmpl, input)

	http.HandleFunc("/cyoa/", handler.HandleRenderStory())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	log.Println("listening to port :8080")

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
