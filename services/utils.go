package services

import (
	"errors"
	"math"
	"strconv"
	"time"
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

func getDayFromDate(s string) (int, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, s)
	if err != nil {
		return 0, errors.New("Date format is invalid")
	}

	return t.Day(), nil
}

func isBetweenTwoAndFour(s string) (bool, error) {
	layout := "15:04"
	t, err := time.Parse(layout, s)
	if err != nil {
		return false, errors.New("Time format is invalid")
	}

	startTime, _ := time.Parse(layout, "14:00")
	endTime, _ := time.Parse(layout, "16:00")

	return t.After(startTime) && t.Before(endTime), nil
}
