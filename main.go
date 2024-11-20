package main

import (
	"receipt-processor/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterReceiptRoutes(router)

	router.Run("localhost:8080")
}
