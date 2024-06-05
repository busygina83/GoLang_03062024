package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID      uint
	Name    string
	Email   string
	GroupID uint
	Group   Group
}

type Group struct {
	ID   uint
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("/tmp/example02.db"), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Group{}, &User{})
	if err != nil {
		panic(err)
	}

	g := Group{Name: "The Ones"}
	u := User{Name: "User", Email: "user@gmail.com", Group: g}
	db.Create(&u)
	fmt.Println("Created", u)

	var recovered User
	db.Preload("Group").First(&recovered, 1)
	fmt.Println("Recovered", recovered)
}