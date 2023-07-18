package models

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
)

func AddNewComment(userId int, ticketId int, text string) error {

	var comment entity.Comment = entity.Comment{
		Text:     text,
		UserId:   userId,
		TicketId: ticketId,
	}

	initilizers.DB.Create(&comment)

	return nil
}
