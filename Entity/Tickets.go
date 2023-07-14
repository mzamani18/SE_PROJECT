package entity

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	UnitPrice     int64
	Capacity      int64
	FirstCapacity int64
	CompanyName   string
	From          string
	To            string
	Time          time.Time
}

type TicketData struct {
	UnitPrice     int64
	Capacity      int64
	FirstCapacity int64
	CompanyName   string
	From          string
	To            string
	Time          string
}
