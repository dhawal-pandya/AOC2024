package main

import (
	"regexp"
	"sort"
	"strconv"
)

// Function to calculate the sum of valid, enabled multiplications
func CalculateSumWithConditions(input string) int {
	// Regex for valid mul(X,Y), do(), and don't()
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	// Initialize variables
	sum := 0
	mulEnabled := true // Multiplications are enabled by default

	// Tokenize input based on recognized patterns
	tokens := mulRe.FindAllStringSubmatchIndex(input, -1)
	tokens = append(tokens, doRe.FindAllStringIndex(input, -1)...)
	tokens = append(tokens, dontRe.FindAllStringIndex(input, -1)...)

	// Sort tokens by their position in the string
	sortedTokens := make([][]int, len(tokens))
	copy(sortedTokens, tokens)
	sort.Slice(sortedTokens, func(i, j int) bool {
		return sortedTokens[i][0] < sortedTokens[j][0]
	})

	// Process each token
	for _, token := range sortedTokens {
		subStr := input[token[0]:token[1]]

		if mulRe.MatchString(subStr) && mulEnabled {
			// Match mul(X, Y) and calculate product if enabled
			matches := mulRe.FindStringSubmatch(subStr)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			sum += x * y
		} else if doRe.MatchString(subStr) {
			// Enable mul instructions
			mulEnabled = true
		} else if dontRe.MatchString(subStr) {
			// Disable mul instructions
			mulEnabled = false
		}
	}

	return sum
}
