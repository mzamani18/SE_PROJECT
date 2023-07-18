package controller

import (
	entity "TikOn/Entity"
	models "TikOn/Models"
	view "TikOn/View"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	userId, isExist := c.Get("user")
	ticketId, err := strconv.Atoi(c.Param("ticket_id"))

	if !isExist || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please set login and set cookie",
		})
		return
	}

	type reqBody struct {
		Text string
	}

	var text reqBody
	if c.Bind(&text) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please enter voucher code",
		})
		return
	}

	err = models.AddNewComment(int(userId.(entity.User).ID), ticketId, text.Text)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "succesful",
	})
}

func CommentsOfTicket(c *gin.Context) {
	ticketId, err := strconv.Atoi(c.Param("ticket_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "get me ticket id",
		})
		return
	}

	comments := view.CommentView(ticketId)

	c.JSON(http.StatusBadRequest, gin.H{
		"message":  "sucessful",
		"comments": comments,
	})
}
