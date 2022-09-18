package functions

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Database() {
	db, openErr := sql.Open("sqlite3", "EveryNasa.db")
	if openErr != nil {
		panic(openErr)
	}

	var exists bool
	QueryErr := db.QueryRow("SELECT EXISTS(SELECT name FROM sqlite_master WHERE type='table' AND name='settings')").Scan(&exists)
	if QueryErr != nil {
		panic(QueryErr)
	}

	if exists == false {
		sqlTable := `
			CREATE TABLE IF NOT EXISTS settings (
				lang TEXT DEFAULT 'en',
				autostart INTEGER DEFAULT 0,
				autoupdate INTEGER DEFAULT 0
			);`
		_, CreateTableErr := db.Exec(sqlTable)
		if CreateTableErr != nil {
			panic(CreateTableErr)
		}

		stmt, InsertErr := db.Prepare("INSERT INTO settings(lang, autostart, autoupdate) values(?,?,?)")
		if InsertErr != nil {
			panic(InsertErr)
		}

		_, ExecErr := stmt.Exec("en", 0, 0)
		if ExecErr != nil {
			panic(ExecErr)
		}
	}
}
