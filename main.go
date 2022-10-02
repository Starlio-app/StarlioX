package main

import (
	"net/http"

	"github.com/Redume/EveryNasa/api/controllers"
	"github.com/Redume/EveryNasa/functions"
	"github.com/Redume/EveryNasa/web/pages"
	"github.com/getlantern/systray"
	"github.com/gorilla/mux"
)

func main() {
	go functions.Logger("EveryNasa started")
	go functions.Database()
	go systray.Run(functions.Tray, functions.Quit)
	go functions.StartWallpaper()

	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, r)
		})
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	http.HandleFunc("/", page.GalleryHandler)
	http.HandleFunc("/settings", page.SettingsHandler)
	http.HandleFunc("/about", page.AboutHandler)

	router.HandleFunc("/api/get/settings", controllers.SettingsGet).Methods("GET")
	router.HandleFunc("/api/get/version", controllers.Version).Methods("GET")

	router.HandleFunc("/api/update/settings", controllers.SettingsUpdate).Methods("POST")
	router.HandleFunc("/api/update/wallpaper", controllers.WallpaperUpdate).Methods("POST")
	router.HandleFunc("/api/update/startup", controllers.Startup).Methods("POST")
	router.HandleFunc("/api/create/label", controllers.CreateLabel).Methods("POST")

	go func() {
		err := http.ListenAndServe(":4662", nil)
		if err != nil {
			functions.Logger(err.Error())
		}
	}()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		functions.Logger(err.Error())
	}
}
