package entity

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserId          int  `gorm:"unique;not null"`
	User            User `gorm:"references:ID"`
	TotalPrice      int
	PayablePrice    int
	VoucherDiscount int  `gorm:"default:0"`
	CartId          int  `gorm:"unique;not null"`
	Cart            Cart `gorm:"references:ID"`
}

type PaymentData struct {
	UserId          int
	TotalPrice      int
	PayablePrice    int
	VoucherDiscount int
	CartId          int
}
