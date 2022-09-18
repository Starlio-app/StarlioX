package page

import "net/http"

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/src/about.html")
}
