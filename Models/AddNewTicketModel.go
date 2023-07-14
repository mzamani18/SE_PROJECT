package models

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	"errors"
	"fmt"
	"time"
)

func AddNewTicket(ticketData *entity.TicketData, ticketTime time.Time) (*entity.Ticket, error) {

	fmt.Println(ticketData.Time)

	ticket := entity.Ticket{
		UnitPrice:     ticketData.UnitPrice,
		Capacity:      ticketData.Capacity,
		FirstCapacity: ticketData.FirstCapacity,
		CompanyName:   ticketData.CompanyName,
		From:          ticketData.From,
		To:            ticketData.To,
		Time:          ticketTime,
	}

	result := initilizers.DB.Create(&ticket)

	if result.Error != nil {
		return nil, errors.New("can't add new user to DB")
	}

	return &ticket, nil
}
