package entity

import "time"

type CartItem struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Price     int
	CartId    int
	Cart      Cart `gorm:"references:ID"`
	TicketId  int
	Ticket    Ticket `gorm:"references:ID"`
	UserId    int
	User      User `gorm:"references:ID"`
}

type CartItemData struct {
	Price    int
	TicketId int
	UserId   int
	CartId   int
}
