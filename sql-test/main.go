package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "/home/sen0324/testdbs/gee.db")
	defer func() { _ = db.Close() }()

	row := db.QueryRow("SELECT * FROM user LIMIT 1")
	var name string
	var age int
	err := row.Scan(&name, &age)
	if err == nil {
		fmt.Println(name, age)
	} else {
		fmt.Print(err.Error())
	}
}
