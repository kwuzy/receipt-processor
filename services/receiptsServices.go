package services

import (
	"errors"
	"fmt"
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

func GetReceiptPoints(id string) (int, error) {
	receipt, exists := GetReceiptByID(id)
	if !exists {
		return 0, errors.New("Receipt not found")
	}

	points := 0
	// One point for every alphanumeric character in the retailer name.
	points += countAlphanumeric(receipt.Retailer)
	// 50 points if the total is a round dollar amount with no cents.
	totalPrice, err := convertStringToPrice(receipt.Total)
	if err != nil {
		return 0, err
	}
	fmt.Print(totalPrice)
	// 25 points if the total is a multiple of 0.25.
	// 5 points for every two items on the receipt.
	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	// 6 points if the day in the purchase date is odd.
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.

	return points, nil
}
