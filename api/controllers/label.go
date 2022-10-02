package controllers

import (
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
)

var CreateLabel = func(w http.ResponseWriter, r *http.Request) {
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

	utils.Respond(w, utils.Message(true, "The shortcut was successfully created"))
}
