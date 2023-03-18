package utils

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Database() {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		Logger(err.Error())
	}

	if !TableExists(db, "settings") {
		sqlTable := `
			CREATE TABLE IF NOT EXISTS settings (
			    startup INTEGER DEFAULT 0,
			    wallpaper INTEGER DEFAULT 0,
			    save_logg INTEGER DEFAULT 1	                                    
			);`

		_, err = db.Exec(sqlTable)
		if err != nil {
			Logger(err.Error())
		}

		stmt, err := db.Prepare("INSERT INTO settings(startup, wallpaper, save_logg) values(?,?,?)")
		if err != nil {
			Logger(err.Error())
		}

		_, err = stmt.Exec(0, 0, 1)
		if err != nil {
			Logger(err.Error())
		}
	}

	if !TableExists(db, "favorite") {
		sqlTable := `
			CREATE TABLE IF NOT EXISTS favorite (
			    title TEXT,
			    explanation TEXT,
			    copyright TEXT,
			    date TEXT,
			    url TEXT,
			    hdurl TEXT,
			    media_type TEXT
			);`

		_, err = db.Exec(sqlTable)
		if err != nil {
			Logger(err.Error())
		}
	}
}

func GetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "EveryNasa.db")
	if err != nil {
		Logger(err.Error())
	}

	return db
}

func TableExists(db *sql.DB, name string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT name FROM sqlite_master WHERE type='table' AND name=?)", name).Scan(&exists)
	if err != nil {
		Logger(err.Error())
	}

	return exists
}
