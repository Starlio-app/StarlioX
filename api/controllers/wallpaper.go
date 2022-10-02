package controllers

import (
	"net/http"

	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
	"github.com/reujab/wallpaper"
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
		utils.Respond(w, utils.Message(false, "An error occurred while setting the wallpaper"))
		return
	}

	utils.Respond(w, utils.Message(true, "Wallpaper successfully updated"))
}
