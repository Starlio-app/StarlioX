package controllers

import (
	"gopkg.in/yaml.v2"
	"net/http"
	"os"

	"github.com/Redume/EveryNasa/api/utils"
)

var Version = func(w http.ResponseWriter, r *http.Request) {
	file, readErr := os.ReadFile("config.yaml")
	if readErr != nil {
		panic(readErr)
	}

	data := make(map[interface{}]interface{})

	marshalErr := yaml.Unmarshal(file, &data)
	if marshalErr != nil {
		panic(marshalErr)
	}

	utils.Respond(w, utils.Message(true, data["version"].(string)))
}
