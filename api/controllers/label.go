package controllers

import (
	"os"
	"os/user"
	"strings"

	"github.com/Redume/Starlio/utils"
	"github.com/gofiber/fiber/v2"
)

var CreateLabel = func(c *fiber.Ctx) error {
	u, err := user.Current()
	if err != nil {
		utils.Logger(err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		utils.Logger(err.Error())
	}

	dir = strings.Replace(dir, "\\", "\\\\", -1) + "\\Starlio.exe"

	err = utils.CreateLnk(dir, strings.Replace(u.HomeDir, "\\", "\\\\", -1)+"\\Desktop\\Starlio.lnk")
	if err != nil {
		utils.Logger(err.Error())
	}

	utils.Respond(c, utils.Message(true, "The shortcut was successfully created"))
	return nil
}
