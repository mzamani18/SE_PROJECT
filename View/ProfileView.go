package view

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
)

func GetUserTickets(userId int) []entity.Ticket {
	var user entity.User
	initilizers.DB.First(&user, userId)

	var cartItems []entity.CartItem

	initilizers.DB.Where("user_id= ? AND cart_status = ?", user.ID, entity.CLOSED).Find(&cartItems)

	var tickets []entity.Ticket

	for i := 0; i < len(cartItems); i++ {
		tickets = append(tickets, cartItems[i].Ticket)
	}

	return tickets
}
