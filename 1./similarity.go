package main

// CalculateSimilarityScore calculates the similarity score between two lists
func CalculateSimilarityScore(left, right []int) int {
	// create a map to count occurrences of numbers in the right list
	rightCount := make(map[int]int)
	for _, num := range right {
		rightCount[num]++
	}

	similarityScore := 0
	for _, num := range left {
		similarityScore += num * rightCount[num]
	}

	return similarityScore
}
