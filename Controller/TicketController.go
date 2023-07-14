package controller

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	models "TikOn/Models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddTicket(c *gin.Context) {
	var requestBody entity.TicketData

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_msg": "bad request",
		})
		return
	}

	timeOfTicket, err := time.Parse("2006-01-02 15:04", requestBody.Time)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_msg": "invalid time format",
		})
		return
	}

	ticket, err := models.AddNewTicket(&requestBody, timeOfTicket)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_msg": "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"ticket":  ticket,
	})
}

func TicketAction(c *gin.Context) {
	var tickets []entity.Ticket
	initilizers.DB.Find(&tickets)

	c.JSON(http.StatusOK, gin.H{
		"tickets": tickets,
	})
}

func SearchAction(c *gin.Context) {
	var tickets []entity.Ticket

	type body struct {
		From string
		To   string
	}
	var requestBody body

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request body",
		})
	}

	initilizers.DB.Where("`from` like ? And `to` like ?", "%"+requestBody.From+"%", "%"+requestBody.To+"%").Find(&tickets)

	c.JSON(http.StatusOK, gin.H{
		"tickets": tickets,
	})
}
