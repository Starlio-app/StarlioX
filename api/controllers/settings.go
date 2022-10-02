package controllers

import (
	"database/sql"
	"fmt"
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

	var startup, wallpaper int

	for query.Next() {
		err := query.Scan(&startup, &wallpaper)
		if err != nil {
			functions.Logger(err.Error())
		}
		var data = map[string]interface{}{"startup": startup, "wallpaper": wallpaper}
		utils.Respond(w, data)
	}
}

var SettingsUpdate = func(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		functions.Logger(err.Error())
	}

	wallpaper := r.FormValue("wallpaper")
	startup := r.FormValue("startup")

	fmt.Println("S: "+startup, "W: "+wallpaper)

	if startup == "" && wallpaper == "" {
		utils.Respond(w, utils.Message(false, "All fields are required."))
		return
	}

	if wallpaper != "" {
		_, err := db.Exec("UPDATE settings SET wallpaper = ?", wallpaper)
		if err != nil {
			functions.Logger(err.Error())
		}

		if wallpaper == "1" {
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
