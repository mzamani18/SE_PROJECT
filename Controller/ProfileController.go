package controller

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersTickets(c *gin.Context) {
	user_id, isExist := c.Get("user")

	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "set your cookie",
		})
		return
	}

	var user entity.User
	initilizers.DB.First(&user, (user_id.(entity.User)).ID)

	var cartItems []entity.CartItem

	initilizers.DB.Where("user_id= ?", user.ID).Find(&cartItems)

	var tickets []entity.Ticket

	for i := 0; i < len(cartItems); i++ {
		tickets = append(tickets, cartItems[i].Ticket)
	}

	c.JSON(http.StatusOK, gin.H{
		"tickets": tickets,
	})
}
