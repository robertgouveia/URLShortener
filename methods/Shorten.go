package methods

import (
	"net/http"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
