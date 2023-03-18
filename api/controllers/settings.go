package controllers

import (
	"database/sql"
	"github.com/Redume/EveryNasa/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

var SettingsGet = func(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		utils.Logger(err.Error())
	}

	querySettings, err := db.Query("SELECT * FROM settings")

	if err != nil {
		utils.Logger(err.Error())
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			utils.Logger(err.Error())
		}
	}(querySettings)

	var startup, wallpaper, save_logg int

	for querySettings.Next() {
		err := querySettings.Scan(&startup, &wallpaper, &save_logg)
		if err != nil {
			utils.Logger(err.Error())
		}

		var data = map[string]interface{}{
			"startup":   startup,
			"wallpaper": wallpaper,
			"save_logg": save_logg,
		}

		utils.Respond(c, data)
	}

	return nil
}

var SettingsUpdate = func(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		utils.Logger(err.Error())
	}

	startup := c.FormValue("startup")
	wallpaper := c.FormValue("wallpaper")
	save_logg := c.FormValue("save_logg")

	if startup == "" && wallpaper == "" && save_logg == "" {
		utils.Respond(c, utils.Message(false, "All fields are required."))
		return nil
	}

	if wallpaper != "" {
		_, err := db.Exec("UPDATE settings SET wallpaper = ?", wallpaper)
		if err != nil {
			utils.Logger(err.Error())
		}

		if wallpaper == "1" {
			go utils.StartWallpaper()
		}
	}

	if startup != "" {
		_, err := db.Exec("UPDATE settings SET startup = ?", startup)
		if err != nil {
			utils.Logger(err.Error())
		}
	}

	if save_logg != "" {
		_, err := db.Exec("UPDATE settings SET save_logg = ?", save_logg)
		if err != nil {
			utils.Logger(err.Error())
		}
	}

	utils.Respond(c, utils.Message(true, "The settings have been applied successfully."))
	return nil
}
