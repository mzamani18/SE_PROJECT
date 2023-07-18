package main

import (
	entity "TikOn/Entity"
	initilizers "TikOn/Initilizers"
)

func init() {
	initilizers.LoadEnvVariables()
	initilizers.ConnectDataBase()
}

func main() {
	initilizers.DB.AutoMigrate(&entity.User{})
	initilizers.DB.AutoMigrate(&entity.Voucher{})
	initilizers.DB.AutoMigrate(&entity.Cart{})
	initilizers.DB.AutoMigrate(&entity.Ticket{})
	initilizers.DB.AutoMigrate(&entity.CartItem{})
	initilizers.DB.AutoMigrate(&entity.Payment{})
	initilizers.DB.AutoMigrate(&entity.Wallet{})
	initilizers.DB.AutoMigrate(&entity.Comment{})
}
