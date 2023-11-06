package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	fmt.Println("listening to port :80")

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
