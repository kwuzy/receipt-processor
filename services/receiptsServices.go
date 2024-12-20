package services

import (
	"errors"
	"math"
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

	// Data Conversions
	totalPrice, err := convertStringToPrice(receipt.Total)
	if err != nil {
		return 0, err
	}
	day, err := getDayFromDate(receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	isBetweenTwoAndFour, err := isBetweenTwoAndFour(receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}

	// Points Calculations
	points := 0
	// One point for every alphanumeric character in the retailer name.
	points += countAlphanumeric(receipt.Retailer)
	// 50 points if the total is a round dollar amount with no cents.
	if totalPrice == float64(int(totalPrice)) {
		points += 50
	}
	// 25 points if the total is a multiple of 0.25.
	if int(totalPrice*100)%25 == 0 {
		points += 25
	}
	// 5 points for every two items on the receipt.
	points += 5 * int(len(receipt.Items)/2)
	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		if isItemDescriptionMultipleOfThree(item.ShortDescription) {
			itemPrice, err := convertStringToPrice(item.Price)
			if err != nil {
				return 0, err
			}
			points += int(math.Ceil(itemPrice * 0.2))
		}
	}
	// 6 points if the day in the purchase date is odd.
	if day%2 == 1 {
		points += 6
	}
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if isBetweenTwoAndFour {
		points += 10
	}

	return points, nil
}
