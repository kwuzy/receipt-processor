package routes

import (
	"net/http"
	"receipt-processor/models"

	"github.com/gin-gonic/gin"
)

func PostReceipt(c *gin.Context) {
	var receipt models.Receipt

	if err := c.BindJSON(&receipt); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": "hello from post"})
}

func RegisterReceiptRoutes(r *gin.Engine) {
	receipts := r.Group("/receipts")
	{
		receipts.POST("/process", PostReceipt)
	}
}
