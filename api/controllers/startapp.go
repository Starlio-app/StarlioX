package controllers

import (
	"github.com/Redume/EveryNasa/api/utils"
	"golang.org/x/sys/windows/registry"
	"net/http"
	"os"
	"path"
	"strings"
)

var AddStartApp = func(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := path.Join(dir, "EveryNasa.exe")

	k, err := registry.OpenKey(registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Run`,
		registry.QUERY_VALUE|registry.SET_VALUE)

	if err != nil {
		panic(err)
	}

	defer k.Close()

	err = k.SetStringValue("EveryNasa", strings.Replace(path, "/", "\\", -1))
	if err != nil {
		panic(err)
	}

	utils.Respond(w, utils.Message(true, "EveryNasa was added to startup apps"))
}

var RemoveStartApp = func(w http.ResponseWriter, r *http.Request) {
	k, err := registry.OpenKey(registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Run`,
		registry.QUERY_VALUE|registry.SET_VALUE)

	if err != nil {
		panic(err)
	}

	defer k.Close()

	err = k.DeleteValue("EveryNasa")
	if err != nil {
		panic(err)
	}

	utils.Respond(w, utils.Message(true, "EveryNasa was removed from startup apps"))
}
