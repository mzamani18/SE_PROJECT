package entity

import (
	"time"

	"gorm.io/gorm"
)

type VoucherStatus string

const (
	NOT_USED     VoucherStatus = "not_used"
	USED         VoucherStatus = "used"
	CART_RESERVE VoucherStatus = "cart_reserve"
)

type Voucher struct {
	gorm.Model
	VoucherDiscount int
	ExpDate         time.Time
	VoucherStatus   VoucherStatus `gorm:"type:enum('not_used', 'used', 'cart_reserve')";"column:CartStatus"`
}

type VoucherData struct {
	VoucherDiscount int
	ExpDate         time.Time
	VoucherStatus   VoucherStatus
}
