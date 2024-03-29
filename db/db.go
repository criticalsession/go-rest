package db

import (
	"database/sql"
	"fmt"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email VARCHAR(255) NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create users table")
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			location VARCHAR(255) NOT NULL,
			date_time DATETIME NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create events table")
	}

	createRegistartionsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (event_id) REFERENCES events(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createRegistartionsTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create registrations table")
	}
}
