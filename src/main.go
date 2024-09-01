package main

import (
	"URLShortner/methods"
	"URLShortner/redis"
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("html/index.html"))

//panic if template is not found

func render(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		client := redis.GetClient()
		ctx := redis.GetContext()

		fmt.Println(r.URL.Path)
		if url, err := client.Get(ctx, "http://localhost:8080"+r.URL.Path).Result(); err == nil {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			http.Error(w, "404 Not found", http.StatusNotFound)
			return
		}
	}
	data := map[string]string{}

	tmpl.Execute(w, data)
	// you can pass data to the template by passing a map
}

func main() {
	redis.Initialize("redis:6379")
	defer redis.GetClient().Close()
	http.HandleFunc("/", render)
	http.HandleFunc("/shorten", methods.Shorten)
	http.ListenAndServe(":8080", nil)
}
