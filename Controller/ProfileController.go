package controller

import (
	entity "TikOn/Entity"
	view "TikOn/View"
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

	var tickets []entity.Ticket = view.GetUserTickets(int(user_id.(entity.User).ID))

	c.JSON(http.StatusOK, gin.H{
		"tickets": tickets,
	})
}
