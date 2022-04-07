package models

import (
	"packages/db"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Users []User

func Migrate() {
	db.Database().AutoMigrate(&User{})
}
