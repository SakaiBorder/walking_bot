package main

import (
	"app/db/external"
	"app/db/model"

	"log"
)

func main() {
	db := external.Connect()
	db.AutoMigrate(&model.User{})

	db.Create(&model.User{LineID: "test"})
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
	}
	defer sqlDB.Close()
}
