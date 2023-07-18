package view

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
)

func CommentView(ticketId int) []entity.Comment {
	var comments []entity.Comment
	initilizers.DB.Where("ticket_id = ?", ticketId).Find(&comments)

	return comments
}
