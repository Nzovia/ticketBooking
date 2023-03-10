package utils

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

  var DB * gorm.DB

func ConnectingToDatabase(){
	var err error

	//dsn := "root:root tcp(127.0.0.1:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(127.0.0.1:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if(err != nil){
		log.Fatal("Failed to connect to the database")
	}
}