package routes

import (
	"net/http"
	"receipt-processor/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostReceipt(c *gin.Context) {
	var receipt models.Receipt

	if err := c.BindJSON(&receipt); err != nil {
		return
	}

	newID := uuid.New().String()

	response := models.ReceiptProcessResponse{
		ID: newID,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterReceiptRoutes(r *gin.Engine) {
	receipts := r.Group("/receipts")
	{
		receipts.POST("/process", PostReceipt)
	}
}
