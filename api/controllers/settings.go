package controllers

import (
	"database/sql"
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

	defer query.Close()

	var lang string
	var autostart int
	var autoupdate int

	for query.Next() {
		err := query.Scan(&lang, &autostart, &autoupdate)
		if err != nil {
			panic(err)
		}
		var data = map[string]interface{}{"lang": lang, "autostart": autostart, "autoupdate": autoupdate}
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

	if lang == "" && autostart == "" && autoupdate == "" {
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

	utils.Respond(w, utils.Message(true, "Settings updated"))
}
