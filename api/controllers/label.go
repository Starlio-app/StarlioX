package controllers

import (
	"os"
	"os/user"
	"strings"

	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
	"github.com/gofiber/fiber/v2"
)

var CreateLabel = func(c *fiber.Ctx) error {
	u, err := user.Current()
	if err != nil {
		functions.Logger(err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		functions.Logger(err.Error())
	}

	dir = strings.Replace(dir, "\\", "\\\\", -1) + "\\EveryNasa.exe"

	err = functions.CreateLnk(dir, strings.Replace(u.HomeDir, "\\", "\\\\", -1)+"\\Desktop\\EveryNasa.lnk")
	if err != nil {
		functions.Logger(err.Error())
	}

	utils.Respond(c, utils.Message(true, "The shortcut was successfully created"))
	return nil
}
