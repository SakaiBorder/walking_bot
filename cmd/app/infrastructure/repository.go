package infrastructure

import (
	"app/db/external"
	"app/db/model"
	"log"
)

func FindByLineId(lineId string) model.User {

	db := external.Connect()
	var user model.User
	db.Where("line_id = ?", lineId).Find(&user)
	log.Println(user)
	return user
}

func Save(user model.User) {
	db := external.Connect()
	db.Save(&user)
}
