package controllers

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"os/user"
	"strings"

	"github.com/Redume/Starlio/utils"
)

var Startup = func(c *fiber.Ctx) error {
	startup := c.FormValue("startup")
	if startup == "" {
		utils.Respond(c, utils.Message(false, "All fields are required."))
		return nil
	}

	if startup == "1" {
		SetStartup(c)
	} else if startup == "0" {
		RemoveStartup(c)
	} else {
		utils.Respond(c, utils.Message(false, "Invalid field."))
		return nil
	}

	return nil
}

var SetStartup = func(c *fiber.Ctx) error {
	u, err := user.Current()
	if err != nil {
		utils.Logger(err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		utils.Logger(err.Error())
	}

	dir = strings.Replace(dir, "\\", "\\\\", -1) + "\\Starlio.exe"

	err = utils.CreateLnk(dir, strings.Replace(u.HomeDir, "\\", "\\\\", -1)+"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\Starlio.lnk")
	if err != nil {
		utils.Logger(err.Error())
	}

	utils.Respond(c, utils.Message(true, "The settings have been applied successfully."))
	return nil
}

var RemoveStartup = func(c *fiber.Ctx) error {
	u, err := user.Current()
	if err != nil {
		utils.Logger(err.Error())
	}

	err = os.Remove(strings.Replace(u.HomeDir, "\\", "\\\\", -1) + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\Starlio.lnk")
	if err != nil {
		utils.Logger(err.Error())
	}

	utils.Respond(c, utils.Message(true, "The settings have been applied successfully."))
	return nil
}
