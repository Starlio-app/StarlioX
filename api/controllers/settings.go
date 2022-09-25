package controllers

import (
	"database/sql"
	"net/http"

	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
	_ "github.com/mattn/go-sqlite3"
)

var SettingsGet = func(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		functions.Logger(err.Error())
	}

	query, err := db.Query("SELECT * FROM settings")
	if err != nil {
		functions.Logger(err.Error())
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			functions.Logger(err.Error())
		}
	}(query)

	var startup, autochangewallpaper int

	for query.Next() {
		err := query.Scan(&startup, &autochangewallpaper)
		if err != nil {
			functions.Logger(err.Error())
		}
		var data = map[string]interface{}{"startup": startup, "autochangewallpaper": autochangewallpaper}
		utils.Respond(w, data)
	}
}

var SettingsUpdate = func(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		functions.Logger(err.Error())
	}

	autochangewallpaper := r.FormValue("autochangewallpaper")
	startup := r.FormValue("startup")

	if startup == "" && autochangewallpaper == "" {
		utils.Respond(w, utils.Message(false, "All fields are required."))
		return
	}

	if autochangewallpaper != "" {
		_, err := db.Exec("UPDATE settings SET autochangewallpaper = ?", autochangewallpaper)
		if err != nil {
			functions.Logger(err.Error())
		}

		if autochangewallpaper == "1" {
			go functions.StartWallpaper()
		}
	}

	if startup != "" {
		_, err := db.Exec("UPDATE settings SET startup = ?", startup)
		if err != nil {
			functions.Logger(err.Error())
		}
	}

	utils.Respond(w, utils.Message(true, "The settings have been applied successfully."))
}
