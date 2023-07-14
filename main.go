package main

import (
	controller "TikOn/Controller"
	initilizers "TikOn/Initilizers"
	middleware "TikOn/MiddleWare"

	"github.com/gin-gonic/gin"
)

func init() {
	initilizers.LoadEnvVariables()
	initilizers.ConnectDataBase()
	initilizers.NewBCryptPasswordEncoder()
}

func main() {
	router := gin.Default()

	// user actions:
	router.POST("/v1/user/sign-up", controller.SignUP)             // Done // test passed
	router.POST("/v1/user/log-in", controller.LogIn)               // Done // test passed
	router.POST("/v1/user/change-pass", controller.ChangePassword) // Done // test passed

	// cart actions
	router.GET("/v1/cart", middleware.RequireAuth, controller.CartAction)                          // Done // test passed
	router.GET("/v1/cart/add/:ticket_id", middleware.RequireAuth, controller.AddToCart)            // Done // test passed
	router.GET("/v1/cart/remove/:cart_item_id", middleware.RequireAuth, controller.RemoveFromCart) // Done // test passed

	// payment actions
	router.POST("/v1/payment", middleware.RequireAuth, controller.PaymentAction)
	router.POST("/v1/payment/deposit", middleware.RequireAuth, controller.Deposit) // Done // test passed
	// router.POST("/v1/voucher/add", middleware.RequireAuth)
	// router.POST("/v1/voucher/remove", middleware.RequireAuth)

	// home page actiions
	router.POST("/v1/ticket/add", middleware.RequireAuth, controller.AddTicket) // Done // test passed
	router.GET("/v1/tickets", controller.TicketAction)                          // Done // test passed
	router.POST("/v1/search", middleware.RequireAuth, controller.SearchAction)  // Done // test passed

	// profile actions
	router.GET("/v1/profile/get-tickets", middleware.RequireAuth, controller.UsersTickets) // Done // test passed

	router.Run()
}
