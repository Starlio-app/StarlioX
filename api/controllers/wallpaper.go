package controllers

import (
	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
	"github.com/gofiber/fiber/v2"
	"github.com/reujab/wallpaper"
)

var WallpaperUpdate = func(c *fiber.Ctx) error {
	var url string
	url = c.FormValue("url")
	if url == "" {
		utils.Respond(c, utils.Message(false, "URL is required"))
		return nil
	}

	err := wallpaper.SetFromURL(url)
	if err != nil {
		functions.Logger(err.Error())
		utils.Respond(c, utils.Message(false, "An error occurred while setting the wallpaper"))
		return nil
	}

	utils.Respond(c, utils.Message(true, "Wallpaper successfully updated"))

	return nil
}
