/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"database/sql"
	"main/cmd"
	"main/db"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() {

	var err error

	db.Con, err = sql.Open("sqlite3", "C:\\Users\\User\\Desktop\\gotest\\db.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Con.Exec(`CREATE TABLE IF NOT EXISTS notes (
						id INTEGER PRIMARY KEY, 
						title TEXT NOT NULL, 
						completed BOOLEAN NOT NULL DEFAULT 0, 
						created_at DATETIME NOT NULL, 
						completed_at DATETIME
					)`)

	if err != nil {
		panic(err)
	}
}

func main() {
	initDB()
	cmd.Execute()
}
