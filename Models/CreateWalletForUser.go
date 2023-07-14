package models

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
)

func CreateWlletForUser(amount int, userId int) {
	wallet := entity.Wallet{
		UserId:   userId,
		Ballance: amount,
	}

	initilizers.DB.Create(&wallet)
}
