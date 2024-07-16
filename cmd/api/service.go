package api

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Function to calculate points based on receipt rules
func calculatePoints(receipt Receipt) int {
	points := 0

	// 1. One point for every alphanumeric character in the retailer name.
	reg := regexp.MustCompile("[^a-zA-Z0-9]")
	cleanRetailer := reg.ReplaceAllString(receipt.Retailer, "") //removes all non alpha numeric characters
	points += len(strings.ReplaceAll(cleanRetailer, " ", ""))   //removes all spaces
	fmt.Println("Points after rule 1:", len(strings.ReplaceAll(cleanRetailer, " ", "")))

	// 2. 50 points if the total is a round dollar amount with no cents.
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && total == math.Floor(total) {
		points += 50
		fmt.Println("Points after rule 2:", 50)
	}

	// 3. 25 points if the total is a multiple of 0.25.
	if total > 0 && math.Mod(total, 0.25) == 0 {
		points += 25
		fmt.Println("Points after rule 3:", 25)
	}

	// 4. 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5
	fmt.Println("Points after rule 4:", (len(receipt.Items)/2)*5)

	// 5. If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength > 0 && trimmedLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
			fmt.Println("Points added in rule 5:", int(math.Ceil(price*0.2)))

		}
	}
	// fmt.Println("Points after rule 5:", points)

	// 6. 6 points if the day in the purchase date is odd.
	purchaseDateTime, err := time.Parse("2006-01-02 15:04", receipt.PurchaseDate+" "+receipt.PurchaseTime)
	if err == nil && purchaseDateTime.Day()%2 != 0 {
		points += 6
		fmt.Println("Points after rule 6:", 6)
	}

	// 7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if err == nil && purchaseDateTime.Hour() >= 14 && purchaseDateTime.Hour() < 16 {
		points += 10
		fmt.Println("Points after rule 7:", 10)
	}

	return points
}
