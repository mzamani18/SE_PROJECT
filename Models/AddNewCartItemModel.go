package models

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	"errors"
)

func AddNewCartItem(cart *entity.Cart, ticketId int, userId int) error {

	var ticket entity.Ticket
	initilizers.DB.First(&ticket, ticketId)

	cartItem := entity.CartItem{
		Price:    int(ticket.UnitPrice),
		TicketId: ticketId,
		UserId:   userId,
		CartId:   int(cart.ID),
	}

	result := initilizers.DB.Create(&cartItem)

	if result.Error != nil {
		return errors.New("")
	}

	initilizers.DB.Model(&cart).Update("payable_price", cart.PayablePrice+cartItem.Price)

	return nil
}
