package view

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	"errors"
)

func CartView(userId int) (entity.Cart, []entity.CartItem, error) {
	var user entity.User
	initilizers.DB.First(&user, userId)

	var cart entity.Cart

	initilizers.DB.Where("user_id = ? AND cart_status = ?", user.ID, entity.OPEN).First(&cart)

	if cart.ID == 0 {
		err := errors.New("your cart is empty, please add ticket to your cart")
		return cart, nil, err
	}

	var cartItems []entity.CartItem

	initilizers.DB.Where("cart_id = ?", cart.ID).Find(&cartItems)

	return cart, cartItems, nil
}
