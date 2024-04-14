package handlers

import (
	"html/template"
	"net/http"
)

type Movies struct {
	Title string
	Year  string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

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

	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.ExecuteTemplate(w, "movies-list-element", Movies{title, year})

}
