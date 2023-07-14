package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	LoginPhone string `gorm:"unique"`
	Mail       string `gorm:"unique"`
	Password   string
	NationalId string `gorm:"unique"`
	Age        int
}

type UserData struct {
	Name       string
	LoginPhone string
	Mail       string
	Password   string
	NationalId string
	Age        int
}
