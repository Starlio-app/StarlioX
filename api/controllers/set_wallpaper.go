package controllers

import (
	"github.com/Redume/EveryNASA/api/utils"
	utils2 "github.com/Redume/EveryNASA/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/reujab/wallpaper"
)

var SetWallpaper = func(c *fiber.Ctx) error {
	url := c.FormValue("url")
	if url == "" {
		utils.Respond(c, utils.Message(false, "URL`s required."))
		return nil
	}

	err := wallpaper.SetFromURL(url)
	if err != nil {
		utils2.Log(err.Error())
		utils.Respond(c, utils.Message(false, "An error occurred while setting the wallpaper."))
		return err
	}

	utils.Respond(c, utils.Message(true, "The wallpaper was successfully installed."))
	return nil
}
