package main

import (
	"assignment2/controllers/items"
	"assignment2/controllers/orders"
	"assignment2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	const PORT = ":8080"
	router := gin.Default()
	db := database.GetConnection()
	
	items.RegisterRoutes(router,db)
	orders.RegisterRoutes(router,db)

	router.Run(PORT)
}