package services

import (
	"errors"
	"math"
	"strconv"
	"unicode"
)

func countAlphanumeric(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			count++
		}
	}
	return count
}

func convertStringToPrice(s string) (float32, error) {
	price, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, errors.New("Price(s) are invalid")
	}

	price = math.Floor(price*100) / 100

	return float32(price), nil
}
