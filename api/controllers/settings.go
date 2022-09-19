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

	var lang string
	var autostart, autoupdate, autochangewallpaper int

	for query.Next() {
		err := query.Scan(&lang, &autostart, &autoupdate, &autochangewallpaper)
		if err != nil {
			panic(err)
		}
		var data = map[string]interface{}{"lang": lang, "autostart": autostart, "autoupdate": autoupdate, "autochangewallpaper": autochangewallpaper}
		utils.Respond(w, data)
	}
}

var SettingsUpdate = func(w http.ResponseWriter, r *http.Request) {
	db, errOpen := sql.Open("sqlite3", "EveryNasa.db")
	if errOpen != nil {
		panic(errOpen)
	}

	lang := r.FormValue("lang")
	autostart := r.FormValue("autostart")
	autoupdate := r.FormValue("autoupdate")
	autochangewallpaper := r.FormValue("autochangewallpaper")

	if lang == "" && autostart == "" && autoupdate == "" && autochangewallpaper == "" {
		utils.Respond(w, utils.Message(false, "All fields are required"))
		return
	}

	if lang != "" {
		_, err := db.Exec("UPDATE settings SET lang = ?", lang)
		if err != nil {
			panic(err)
		}
	}

	if autostart != "" {
		_, err := db.Exec("UPDATE settings SET autostart = ?", autostart)
		if err != nil {
			panic(err)
		}
	}

	if autoupdate != "" {
		_, err := db.Exec("UPDATE settings SET autoupdate = ?", autoupdate)
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
