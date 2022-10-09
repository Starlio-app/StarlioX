package page

import (
	"github.com/Redume/EveryNasa/functions"
	"net/http"
	"strings"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	con := functions.Connected()
	if con == false {
		http.ServeFile(w, r, "web/src/errors/504.html")
		return
	}

	userAgent := r.UserAgent()
	if strings.Contains(userAgent, "Mobile") {
		http.ServeFile(w, r, "web/src/errors/mobile.html")
		return
	}

	if r.URL.Path != "/settings" {
		http.ServeFile(w, r, "web/src/errors/404.html")
		return
	}

	http.ServeFile(w, r, "web/src/settings.html")
}
