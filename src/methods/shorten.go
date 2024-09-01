package methods

import (
	"URLShortner/redis"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("html/short.html"))

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		var data = map[string]string{
			"error": "Invalid URL",
		}

		tmpl.Execute(w, data)
		return
	}

	//client := redis.GetClient()
	//ctx := redis.GetContext()

	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	shortUrl := "http://localhost:8080/" + base64.URLEncoding.EncodeToString(bytes)

	client := redis.GetClient()
	ctx := redis.GetContext()

	err := client.Set(ctx, shortUrl, url, 0).Err()
	if err != nil {
		fmt.Println("set: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = client.Get(ctx, shortUrl).Result()
	if err != nil {
		fmt.Println("get: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println(shortUrl)
	var data = map[string]string{
		"url": shortUrl,
	}
	tmpl.Execute(w, data)
}
