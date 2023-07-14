package models

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	"errors"
)

func UpdateUserPassword(loginPhone string, password string, newPassword string) error {
	var user entity.User

	initilizers.DB.First(&user, "login_phone = ?", loginPhone)

	isValid, err := initilizers.PasswordEncoder.Matches(password, user.Password)

	if err != nil || !isValid {
		return errors.New("invalid password")
	}

	newPassword, err = initilizers.PasswordEncoder.Encode(newPassword)

	if err != nil {
		return errors.New("invalid new password")
	}

	initilizers.DB.Model(&user).Where("id = ?", user.ID).Update("Password", newPassword)
	return nil
}
