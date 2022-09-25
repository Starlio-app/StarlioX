package controllers

import (
	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
	"github.com/reujab/wallpaper"
	"net/http"
)

var WallpaperUpdate = func(w http.ResponseWriter, r *http.Request) {
	var url string
	url = r.FormValue("url")
	if url == "" {
		utils.Respond(w, utils.Message(false, "URL is required"))
		return
	}

	err := wallpaper.SetFromURL(url)
	if err != nil {
		functions.Logger(err.Error())
	}
}
