package main

import (
	"fmt"
	"gophercises/choose-your-own-adventure/internal/handler"
	"gophercises/choose-your-own-adventure/internal/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Change the current working directory to be consistent, regardless of where the program is run
	data, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatalln(err)
	}

	input, err := models.NewAdventureInput(data)
	if err != nil {
		log.Fatalln(err)
	}

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	layoutFiles, err := filepath.Glob("templates/layout/" + "*.html")
	if err != nil {
		log.Fatalln(err)
	}

	pageFiles, err := filepath.Glob("templates/" + "*.html")
	if err != nil {
		log.Fatalln(err)
	}

	templates := make(map[string]*template.Template)

	for _, file := range pageFiles {
		filename := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[filename] = template.Must(template.ParseFiles(files...))
	}

	handler := handler.NewHandler(templates, input)

	http.HandleFunc("/cyoa/", handler.HandleRenderStory())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	log.Println("listening to port :8080")

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
