package controllers

import (
	"database/sql"
	"github.com/Redume/EveryNasa/functions"
	"net/http"

	"github.com/Redume/EveryNasa/api/utils"
	_ "github.com/mattn/go-sqlite3"
)

var SettingsGet = func(w http.ResponseWriter, r *http.Request) {
	db, errOpen := sql.Open("sqlite3", "EveryNasa.db")
	if errOpen != nil {
		panic(errOpen)
	}

	query, errQuery := db.Query("SELECT * FROM settings")
	if errQuery != nil {
		panic(errQuery)
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			panic(err)
		}
	}(query)

	var autostart, autochangewallpaper int

	for query.Next() {
		err := query.Scan(&autostart, &autochangewallpaper)
		if err != nil {
			panic(err)
		}
		var data = map[string]interface{}{"autostart": autostart, "autochangewallpaper": autochangewallpaper}
		utils.Respond(w, data)
	}
}

var SettingsUpdate = func(w http.ResponseWriter, r *http.Request) {
	db, errOpen := sql.Open("sqlite3", "EveryNasa.db")
	if errOpen != nil {
		panic(errOpen)
	}

	autostart := r.FormValue("autostart")
	autochangewallpaper := r.FormValue("autochangewallpaper")

	if autostart == "" && autochangewallpaper == "" {
		utils.Respond(w, utils.Message(false, "All fields are required"))
		return
	}

	if autostart != "" {
		_, err := db.Exec("UPDATE settings SET autostart = ?", autostart)
		if err != nil {
			panic(err)
		}
	}

	if autochangewallpaper != "" {
		_, err := db.Exec("UPDATE settings SET autochangewallpaper = ?", autochangewallpaper)
		if err != nil {
			panic(err)
		}

		if autochangewallpaper == "1" {
			go functions.StartWallpaper()
		}
	}

	utils.Respond(w, utils.Message(true, "Settings updated"))
}
