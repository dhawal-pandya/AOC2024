package main

import (
	"fmt"
)

func main() {
	left, right := ParseIntoArrays(Input)

	totalDistance := CalculateTotalDistance(left, right) // Total Distance: 1319616
	fmt.Println("Total Distance:", totalDistance)

	score := CalculateSimilarityScore(left, right) // Similarity Score: 27267728
	fmt.Println("Similarity Score:", score)
}
