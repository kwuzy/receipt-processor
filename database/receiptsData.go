package database

import (
	"receipt-processor/models"

	"github.com/google/uuid"
)

var receiptStore = make(map[string]models.Receipt)

func ProcessReceipt(receipt models.Receipt) string {
	newID := uuid.New().String()
	receiptStore[newID] = receipt
	return newID
}

func GetReceipt(id string) (models.Receipt, bool) {
	receipt, exists := receiptStore[id]
	return receipt, exists
}
