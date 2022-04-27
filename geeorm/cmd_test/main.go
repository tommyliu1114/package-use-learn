package main

import (
	"fmt"
	"geeorm"

	log "geeorm/log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

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

	var users []User
	s.Find(&users)
	fmt.Printf("users are : %+v \n", users)
	var tUser User
	s.First(&tUser)
	fmt.Printf("tuser is %+v \n", tUser)
}
