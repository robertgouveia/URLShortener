package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))

//panic if template is not found

func render(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{}

	tmpl.Execute(w, data)
	// you can pass data to the template by passing a map
}

func main() {
	http.HandleFunc("/", render)
	http.ListenAndServe(":8080", nil)
}
