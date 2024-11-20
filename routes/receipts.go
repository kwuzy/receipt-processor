package routes

import (
	"net/http"
	"receipt-processor/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func PostReceipt(c *gin.Context) {
	var receipt models.Receipt
	validate := validator.New()

	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := validate.Struct(receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Receipt format"})
		return
	}

	// TODO: validate multiple uploads
	// TODO: total price is equal to sum of item prices
	// TODO: validate date and time

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
