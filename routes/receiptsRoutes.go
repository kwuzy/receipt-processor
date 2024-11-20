package routes

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	// TODO: invalidate multiple uploads of the same receipt
	// TODO: validate total price is equal to sum of item prices
	// TODO: validate date and time

	newID, err := services.ProcessReceipt(receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process receipt"})
		return
	}

	response := models.ReceiptProcessResponse{
		ID: newID,
	}

	c.JSON(http.StatusOK, response)
}

func GetReceiptByID(c *gin.Context) {
	id := c.Param("id")

	receipt, exists := services.GetReceiptByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	c.JSON(http.StatusOK, receipt)
}

func GetReceiptPoints(c *gin.Context) {
	id := c.Param("id")

	points, err := services.GetReceiptPoints(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := models.ReceiptPointsResponse{
		Points: points,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterReceiptRoutes(r *gin.Engine) {
	receipts := r.Group("/receipts")
	{
		receipts.POST("/process", PostReceipt)
		receipts.GET("/:id", GetReceiptByID)
		receipts.GET("/:id/points", GetReceiptPoints)
	}
}
