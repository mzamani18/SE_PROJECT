package controller

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
	models "TikOn/Models"
	view "TikOn/View"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CartAction(c *gin.Context) {
	user_id, isExist := c.Get("user")

	if !isExist {
		c.JSON(http.StatusOK, gin.H{
			"message": "user not found, please log in",
		})
		return
	}
	cart, cartItems, err := view.CartView(int(user_id.(entity.User).ID))

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart_items_data": cartItems,
		"cart_data":       cart,
	})
}

func AddToCart(c *gin.Context) {
	user_id, isExist := c.Get("user")
	ticketId, err := strconv.Atoi(c.Param("ticket_id"))

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "problem",
		})
		return
	}

	var user entity.User
	initilizers.DB.First(&user, (user_id.(entity.User)).ID)

	if !isExist {
		c.JSON(http.StatusOK, gin.H{
			"message": "user not found, please log in",
		})
		return
	}

	var cart entity.Cart

	initilizers.DB.Where("user_id = ? AND cart_status = ?", user.ID, entity.OPEN).First(&cart)

	if cart.ID == 0 {
		cart, err = models.AddNewCart(int(user.ID))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "can't make a cart for user",
			})
			return
		}
	}

	err = models.AddNewCartItem(&cart, int(ticketId), int(user.ID))

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error in add to cat",
		})
		return
	}

	var cartItems []entity.CartItem

	initilizers.DB.Where("cart_id = ?", cart.ID).Find(&cartItems)

	c.JSON(http.StatusOK, gin.H{
		"cart":            cart,
		"cart_items_data": cartItems,
	})
}

func RemoveFromCart(c *gin.Context) {
	user_id, isExist := c.Get("user")
	cartItemId := c.Param("cart_item_id")

	var user entity.User
	initilizers.DB.First(&user, (user_id.(entity.User)).ID)

	if !isExist {
		c.JSON(http.StatusOK, gin.H{
			"message": "user not found, please log in",
		})
		return
	}

	var cartItem entity.CartItem

	initilizers.DB.First(&cartItem, cartItemId)

	var cart entity.Cart

	initilizers.DB.First(&cart, cartItem.CartId)

	if cart.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "your cart is empty, please add ticket to your cart",
		})
		return
	}

	initilizers.DB.Model(&cart).Update("payable_price", cart.PayablePrice-cartItem.Price)
	initilizers.DB.Delete(&cartItem)

	var cartItems []entity.CartItem

	initilizers.DB.Where("cart_id = ?", cart.ID).Find(&cartItems)

	c.JSON(http.StatusOK, gin.H{
		"cart_items_data": cartItems,
		"cart_data":       cart,
	})
}
