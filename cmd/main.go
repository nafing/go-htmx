package main

import (
	"fmt"
	"log"
	"net/http"

	handler "go-htmx/handlers"
)

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/", handler.LoadHome)
	http.HandleFunc("/add-movie", handler.AddMovie)
	http.HandleFunc("/about", handler.About)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
