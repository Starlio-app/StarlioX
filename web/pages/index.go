package page

import (
	"net/http"

	"github.com/Redume/EveryNasa/functions"
)

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
	con := functions.Connected()
	if con == false {
		http.ServeFile(w, r, "web/src/errors/504.html")
		return
	}

	if r.URL.Path != "/" {
		http.ServeFile(w, r, "web/src/errors/404.html")
		return
	}

	http.ServeFile(w, r, "web/src/gallery.html")
}
