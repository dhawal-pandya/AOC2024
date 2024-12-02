package main

import (
	"math"
	"sort"
)

// CalculateTotalDistance calculates the total distance between two lists
func CalculateTotalDistance(left, right []int) int {
	// sort both slices
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := range left {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	return totalDistance
}
