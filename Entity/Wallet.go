package entity

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserId   int  `gorm:"unique;not null"`
	User     User `gorm:"references:ID"`
	Ballance int  `gorm:"default:0"`
}

type WalletData struct {
	UserId   int
	Ballance int
}
