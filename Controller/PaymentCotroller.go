package controller

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PaymentAction(c *gin.Context) {
	// as we don't have a bank service so we assume that we don't have any wallet and
	// if you click on payment we will pay it and payment status is successful.
	user_id, _ := c.Get("user")

	var user entity.User
	initilizers.DB.First(&user, (user_id.(entity.User)).ID)

	type reqBody struct {
		CartId int
	}
	var requestBody reqBody

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "filed to read request body",
		})
	}

	var cart entity.Cart
	initilizers.DB.First(&cart, requestBody.CartId)

	var wallet entity.Wallet
	initilizers.DB.Where("user_id = ?", user.ID).First(&wallet)

	if cart.PayablePrice > wallet.Ballance {
		c.JSON(http.StatusOK, gin.H{
			"message": "insufficient balance",
		})
	}

	var payment entity.PaymentData = entity.PaymentData{
		UserId:          int(user.ID),
		TotalPrice:      cart.PayablePrice,
		PayablePrice:    cart.PayablePrice,
		VoucherDiscount: 0,
		CartId:          int(cart.ID),
	}

	initilizers.DB.Model(&wallet).Update("ballance", wallet.Ballance-cart.PayablePrice)
	initilizers.DB.Model(&cart).Update("cart_status", entity.CLOSED)

	var finalPayment entity.Payment = entity.Payment{
		UserId:          payment.UserId,
		TotalPrice:      payment.TotalPrice,
		PayablePrice:    payment.PayablePrice,
		VoucherDiscount: 0,
		CartId:          payment.CartId,
	}

	initilizers.DB.Create(&finalPayment)

	c.JSON(http.StatusOK, gin.H{
		"message":      "suceesful payment",
		"payment_data": payment,
	})
}

func Deposit(c *gin.Context) {
	user_id, _ := c.Get("user")

	var user entity.User
	initilizers.DB.First(&user, (user_id.(entity.User)).ID)

	type reqBody struct {
		Amount int
	}

	requestBody := reqBody{}

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "faied to read body",
		})
		return
	}

	var wallet entity.Wallet
	initilizers.DB.Where("user_id = ?", user.ID).First(&wallet)
	initilizers.DB.Model(&wallet).Update("ballance", wallet.Ballance+requestBody.Amount)

	c.JSON(http.StatusOK, gin.H{
		"message": "your wallet charged",
	})
}
