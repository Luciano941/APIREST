package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "luciano:1234@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"

var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error al conectar")
		return nil
	} else {
		fmt.Println("Conectado")
		return db
	}
}
