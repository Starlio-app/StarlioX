package functions

import (
	"database/sql"
	"encoding/json"
	"github.com/rodkranz/fetch"

	_ "github.com/mattn/go-sqlite3"
)

func Database() {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		Logger(err.Error())
	}

	var existsSettings bool
	err = db.QueryRow("SELECT EXISTS(SELECT name FROM sqlite_master WHERE type='table' AND name='settings')").Scan(&existsSettings)

	if err != nil {
		Logger(err.Error())
	}

	if existsSettings == false {
		sqlTable := `
			CREATE TABLE IF NOT EXISTS settings (
			    startup INTEGER DEFAULT 0,
			    wallpaper INTEGER DEFAULT 0,
			    save_logg INTEGER DEFAULT 0
			);`

		_, err = db.Exec(sqlTable)
		if err != nil {
			Logger(err.Error())
		}

		stmt, err := db.Prepare("INSERT INTO settings(startup, wallpaper, save_logg) values(?,?,?)")
		if err != nil {
			Logger(err.Error())
		}

		_, err = stmt.Exec(0, 0, 0)
		if err != nil {
			Logger(err.Error())
		}
	}
}

func getDatabase() int {
	client := fetch.NewDefault()
	res, err := client.Get("http://localhost:3000/api/get/settings", nil)
	if err != nil {
		panic(err)
	}

	body, err := res.ToString()
	if err != nil {
		panic(err)
	}

	type DatabaseStruct struct {
		Save_logg int `json:"save_logg"`
	}

	var Database DatabaseStruct
	err = json.Unmarshal([]byte(body), &Database)
	if err != nil {
		panic(err)
	}

	return Database.Save_logg
}
