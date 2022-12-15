package controllers

import (
	"database/sql"
	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

var SettingsGet = func(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		functions.Logger(err.Error())
	}

	querySettings, err := db.Query("SELECT * FROM settings")

	if err != nil {
		functions.Logger(err.Error())
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			functions.Logger(err.Error())
		}
	}(querySettings)

	var startup, wallpaper, save_logg, analytics int

	for querySettings.Next() {
		err := querySettings.Scan(&startup, &wallpaper, &save_logg, &analytics)
		if err != nil {
			functions.Logger(err.Error())
		}

		var data = map[string]interface{}{
			"startup":   startup,
			"wallpaper": wallpaper,
			"save_logg": save_logg,
			"analytics": analytics}

		utils.Respond(c, data)
	}

	return nil
}

var SettingsUpdate = func(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		functions.Logger(err.Error())
	}

	startup := c.FormValue("startup")
	wallpaper := c.FormValue("wallpaper")
	save_logg := c.FormValue("save_logg")
	analytics := c.FormValue("analytics")

	if startup == "" && wallpaper == "" && save_logg == "" && analytics == "" {
		utils.Respond(c, utils.Message(false, "All fields are required."))
		return nil
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

	if save_logg != "" {
		_, err := db.Exec("UPDATE settings SET save_logg = ?", save_logg)
		if err != nil {
			functions.Logger(err.Error())
		}
	}

	if analytics != "" {
		_, err := db.Exec("UPDATE settings SET analytics = ?", analytics)
		if err != nil {
			functions.Logger(err.Error())
		}
	}

	utils.Respond(c, utils.Message(true, "The settings have been applied successfully."))
	return nil
}
