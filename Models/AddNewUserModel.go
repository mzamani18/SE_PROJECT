package models

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	"errors"
)

func AddNewUser(userData *entity.UserData) (int, error) {

	encrptedPass, err := initilizers.PasswordEncoder.Encode(userData.Password)

	if err != nil {
		return 0, errors.New("invalid password")
	}

	user := entity.User{
		Name:       userData.Name,
		LoginPhone: userData.LoginPhone,
		Mail:       userData.Mail,
		Password:   encrptedPass,
		NationalId: userData.NationalId,
		Age:        userData.Age,
	}

	result := initilizers.DB.Create(&user)

	if result.Error != nil {
		return 0, errors.New("can't add new user to DB")
	}

	return int(user.ID), nil
}
