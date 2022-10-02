package controllers

import (
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
)

var Startup = func(w http.ResponseWriter, r *http.Request) {
	startup := r.FormValue("startup")
	if startup == "" {
		utils.Respond(w, utils.Message(false, "All fields are required."))
		return
	}

	if startup == "1" {
		SetStartup(w, r)
	} else if startup == "0" {
		RemoveStartup(w, r)
	} else {
		utils.Respond(w, utils.Message(false, "Invalid field."))
		return
	}
}

var SetStartup = func(w http.ResponseWriter, r *http.Request) {
	u, err := user.Current()
	if err != nil {
		functions.Logger(err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		functions.Logger(err.Error())
	}

	dir = strings.Replace(dir, "\\", "\\\\", -1) + "\\EveryNasa.exe"

	err = functions.CreateLnk(dir, strings.Replace(u.HomeDir, "\\", "\\\\", -1)+"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\EveryNasa.lnk")
	if err != nil {
		functions.Logger(err.Error())
	}

	utils.Respond(w, utils.Message(true, "The settings have been applied successfully."))
}

var RemoveStartup = func(w http.ResponseWriter, r *http.Request) {
	u, err := user.Current()
	if err != nil {
		functions.Logger(err.Error())
	}

	err = os.Remove(strings.Replace(u.HomeDir, "\\", "\\\\", -1) + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\EveryNasa.lnk")
	if err != nil {
		functions.Logger(err.Error())
	}

	utils.Respond(w, utils.Message(true, "The settings have been applied successfully."))
}
