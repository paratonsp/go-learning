package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:Katon123@tcp(localhost:3306)/learning_go"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})

	DB = db
}
