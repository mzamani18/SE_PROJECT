package entity

import "gorm.io/gorm"

type CartStatus string

const (
	OPEN   CartStatus = "open"
	CLOSED CartStatus = "closed"
)

type Cart struct {
	gorm.Model
	PayablePrice    int
	VoucherDiscount int `gorm:"default:0"`
	UserId          int
	User            User       `gorm:"references:ID"`
	VoucherId       int        `gorm:"default:0"`
	Voucher         Voucher    `gorm:"references:ID"`
	CartStatus      CartStatus `gorm:"type:enum('open', 'closed')";"column:VoucherStatus"`
}

type CartData struct {
	PayablePrice    int
	VoucherDiscount int
	UserId          int
	VoucherId       int
	CartStatus      CartStatus
}
