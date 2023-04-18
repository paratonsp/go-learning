package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("paratonsp:Katon#11@tcp(103.54.170.102:3306)/learning_go"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
		fmt.Println(err.Error())
		return
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})

	DB = db
}
