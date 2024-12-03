package main

import (
	"regexp"
	"strconv"
)

// Extracts valid mul instructions and computes their result.
func CalculateSum(input string) int {
	// Regex to match valid mul(X,Y) patterns
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		// Convert X and Y to integers
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}
	return sum
}
