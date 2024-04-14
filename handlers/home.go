package handlers

import (
	"html/template"
	"net/http"
)

type Movies struct {
	Title string
	Year  string
}

var tmpl = template.Must(template.ParseFiles("templates/home.html"))

func LoadHome(w http.ResponseWriter, r *http.Request) {
	// Loads the home.html template
	movies := map[string][]Movies{
		"Movies": {
			{"The Shawshank Redemption", "1994"},
			{"The Godfather", "1972"},
			{"The Dark Knight", "2008"},
		},
	}

	tmpl.Execute(w, movies)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	year := r.FormValue("year")
	// Renders the movies-list-element template
	tmpl.ExecuteTemplate(w, "movies-list-element", Movies{title, year})
}
