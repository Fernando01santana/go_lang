package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", func(response http.ResponseWriter, request *http.Request) {
		templates.ExecuteTemplate(response, "home.html", nil)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
