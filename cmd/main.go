package main

import (
	"fmt"
	"log"
	"net/http"

	home "go-htmx/handlers"
)

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/", home.Home)
	http.HandleFunc("/about", home.About)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
