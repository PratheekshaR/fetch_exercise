package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// Function to calculate points for a receipt
func calculatePoints(receipt Receipt) int {
	totalPoints := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	totalPoints += countAlphanumericChars(receipt.RetailerName)

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	totalAmount, _ := strconv.ParseFloat(receipt.TotalAmount, 64)
	if isRoundDollar(totalAmount) {
		totalPoints += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if isMultipleOf25(totalAmount) {
		totalPoints += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	totalPoints += 5 * (len(receipt.Items) / 2)

	// Rule 5: Apply the rule for each item description being a multiple of 3
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			totalPoints += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	parsedDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if parsedDate.Day()%2 != 0 {
		totalPoints += 6
	}

	// Rule 7: 10 points if the time of purchase is between 2:00pm and 4:00pm
	parsedTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		totalPoints += 10
	}

	return totalPoints
}

// Helper function to count alphanumeric characters in a string
func countAlphanumericChars(s string) int {
	count := 0
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			count++
		}
	}
	return count
}

// Helper function to check if the total is a round dollar amount (no cents)
func isRoundDollar(amount float64) bool {
	return amount == math.Floor(amount)
}

// Helper function to check if the total is a multiple of 0.25
func isMultipleOf25(amount float64) bool {
	return math.Mod(amount, 0.25) == 0
}
