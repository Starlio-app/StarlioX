package controllers

import (
	"net/http"
	"os"

	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
	"gopkg.in/yaml.v2"
)

var Version = func(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		functions.Logger(err.Error())
	}

	data := make(map[interface{}]interface{})

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		functions.Logger(err.Error())
	}

	utils.Respond(w, utils.Message(true, data["version"].(string)))
}
