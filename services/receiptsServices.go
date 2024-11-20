package services

import (
	"receipt-processor/database"
	"receipt-processor/models"
)

func ProcessReceipt(receipt models.Receipt) (string, error) {
	newID := database.ProcessReceipt(receipt)
	return newID, nil
}

func GetReceiptByID(id string) (models.Receipt, bool) {
	receipt, exists := database.GetReceipt(id)
	return receipt, exists
}
