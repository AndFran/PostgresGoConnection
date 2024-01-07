package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		fmt.Printf("Driver %v\n", driver)
	}
}

func openDatbase() (db *sql.DB, err error) {
	db, err = sql.Open("pgx", "postgres://USERNAME:PASSWORD@localhost:5432/THE_DB")
	if err == nil {
		fmt.Println("Connection open")
	}
	return
}
