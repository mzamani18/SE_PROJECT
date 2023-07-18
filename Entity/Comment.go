package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text     string
	TicketId int    `gorm:"default:0"`
	Ticket   Ticket `gorm:"references:ID"`
	UserId   int    `gorm:"default:0"`
	User     User   `gorm:"references:ID"`
}

type CommentData struct {
	Text     string
	TicketId int
	UserId   int
}
