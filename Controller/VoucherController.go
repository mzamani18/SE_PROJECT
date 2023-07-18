package controller

import (
	entity "TikOn/Entity"
	models "TikOn/Models"
	view "TikOn/View"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddVoucherToCart(c *gin.Context) {
	userId, isExist := c.Get("user")

	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please set login and set cookie",
		})
		return
	}

	type reqBody struct {
		VoucherCode string
	}
	var voucher reqBody
	if c.Bind(&voucher) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please enter voucher code",
		})
		return
	}

	err := models.AddVoucher(int(userId.(entity.User).ID), voucher.VoucherCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	cart, _, _ := view.CartView(int(userId.(entity.User).ID))

	c.JSON(http.StatusBadRequest, gin.H{
		"message": cart,
	})
}

func RemoveVoucherFromCart(c *gin.Context) {
	userId, isExist := c.Get("user")

	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please set login and set cookie",
		})
		return
	}

	err := models.RemoveVoucher(int(userId.(entity.User).ID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	cart, _, _ := view.CartView(int(userId.(entity.User).ID))

	c.JSON(http.StatusBadRequest, gin.H{
		"message": cart,
	})
}
