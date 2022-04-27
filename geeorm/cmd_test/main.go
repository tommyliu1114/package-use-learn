package main

import (
	"geeorm"

	log "geeorm/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "/home/sen0324/testdbs/gee.db")
	defer engine.Close()
	s := engine.NewSession()
	row := s.Raw("SELECT * FROM user LIMIT 1").QueryRow()

	var name string
	var age int
	err := row.Scan(&name, &age)
	if err == nil {
		log.Info(name, age)
	} else {
		log.Error(err.Error())
	}

}
