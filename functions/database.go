package functions

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Database() {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		Logger(err.Error())
	}

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT name FROM sqlite_master WHERE type='table' AND name='settings')").Scan(&exists)
	if err != nil {
		Logger(err.Error())
	}

	if exists == false {
		sqlTable := `
			CREATE TABLE IF NOT EXISTS settings (
			    startup INTEGER DEFAULT 0,
			    autochangewallpaper INTEGER DEFAULT 0
			);`

		_, err = db.Exec(sqlTable)
		if err != nil {
			Logger(err.Error())
		}

		stmt, err := db.Prepare("INSERT INTO settings(startup, autochangewallpaper) values(?, ?)")
		if err != nil {
			Logger(err.Error())
		}

		_, err = stmt.Exec(0, 0)
		if err != nil {
			Logger(err.Error())
		}
	}
}
