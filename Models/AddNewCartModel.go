package models

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	"errors"
)

func AddNewCart(userId int) (entity.Cart, error) {

	cart := entity.Cart{
		PayablePrice: 0,
		UserId:       userId,
		// VoucherId:    nil,
		CartStatus: entity.OPEN,
	}

	result := initilizers.DB.Create(&cart)

	if result.Error != nil {
		return cart, errors.New("can't add new user to DB")
	}

	return cart, nil
}
