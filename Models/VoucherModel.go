package models

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	"errors"
	"math"
)

func AddVoucher(userId int, voucherCode string) error {
	var cart entity.Cart
	initilizers.DB.Where("user_id = ? and cart_status = ?", userId, entity.OPEN).First(&cart)

	if cart.ID == 0 {
		return errors.New("yor cart is empty, please add ticket to your cart")
	}

	var voucher entity.Voucher
	initilizers.DB.Where("voucher_code = ? and voucher_status = ? ", voucherCode, entity.NOT_USED).First(&voucher)

	if voucher.ID == 0 {
		return errors.New("invalid voucher code, voucher not found")
	}

	initilizers.DB.Model(&cart).Update("voucher_discount", math.Min(float64(voucher.VoucherDiscount), float64(cart.PayablePrice)))
	initilizers.DB.Model(&cart).Update("payable_price", math.Max(float64(cart.PayablePrice-voucher.VoucherDiscount), 0))
	initilizers.DB.Model(&cart).Update("voucher_id", voucher.ID)

	initilizers.DB.Model(&voucher).Update("voucher_status", entity.USED)

	return nil
}

func RemoveVoucher(userId int) error {
	var cart entity.Cart
	initilizers.DB.Where("user_id = ? and cart_status = ?", userId, entity.OPEN).First(&cart)

	if cart.ID == 0 {
		return errors.New("yor cart is empty, please add ticket to your cart")
	}

	if cart.VoucherId == 0 {
		return errors.New("you dont have voucher on your cart")
	}

	var voucher entity.Voucher
	initilizers.DB.First(&voucher, cart.VoucherId)

	initilizers.DB.Model(&cart).Update("payable_price", cart.PayablePrice+cart.VoucherDiscount)
	initilizers.DB.Model(&cart).Update("voucher_discount", 0)
	initilizers.DB.Model(&cart).Update("voucher_id", 0)

	initilizers.DB.Model(&voucher).Update("voucher_status", entity.NOT_USED)

	return nil
}
