package main

import (
    "math"
    "strconv"
    "strings"
    "time"
)

func calculatePoints(receipt Receipt) int {
    points := 0

    // Rule 1: 1 point for every alphanumeric character in the retailer name.
    points += countAlphanumeric(receipt.Retailer)

    // Rule 2: 50 points if the total is a round dollar amount with no cents.
    if isRoundDollar(receipt.Total) {
        points += 50
    }

    // Rule 3: 25 points if the total is a multiple of 0.25.
    if isMultipleOfQuarter(receipt.Total) {
        points += 25
    }

    // Rule 4: 5 points for every two items on the receipt.
    points += (len(receipt.Items) / 2) * 5

    // Rule 5: Points based on item description length.
    for _, item := range receipt.Items {
        if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
            price, _ := strconv.ParseFloat(item.Price, 64)
            points += int(math.Ceil(price * 0.2))
        }
    }

    // Rule 6: 6 points if the day in the purchase date is odd.
    if isDayOdd(receipt.PurchaseDate) {
        points += 6
    }

    // Rule 7: 10 points if the time of purchase is between 2:00pm and 4:00pm.
    if isBetweenTwoAndFour(receipt.PurchaseTime) {
        points += 10
    }

    return points
}

func countAlphanumeric(s string) int {
    count := 0
    for _, r := range s {
        if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
            count++
        }
    }
    return count
}

func isRoundDollar(total string) bool {
    amount, _ := strconv.ParseFloat(total, 64)
    return amount == float64(int(amount))
}

func isMultipleOfQuarter(total string) bool {
    amount, _ := strconv.ParseFloat(total, 64)
    return math.Mod(amount, 0.25) == 0
}

func isDayOdd(dateStr string) bool {
    date, _ := time.Parse("2006-01-02", dateStr)
    return date.Day()%2 != 0
}

func isBetweenTwoAndFour(timeStr string) bool {
    t, _ := time.Parse("15:04", timeStr)
    return t.Hour() == 14 || (t.Hour() == 15 && t.Minute() == 0)
}
