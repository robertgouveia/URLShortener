package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))

//panic if template is not found

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Message": "Hello, World!",
	}

	tmpl.Execute(w, data)
	// you can pass data to the template by passing a map
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
