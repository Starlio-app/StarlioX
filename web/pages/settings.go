package page

import "net/http"

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/src/settings.html")
}
