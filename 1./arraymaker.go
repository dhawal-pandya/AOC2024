package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseIntoArrays(input string) ([]int, []int) {
	// split the input string into individual numbers
	numbers := strings.Fields(input)

	var arr1, arr2 []int
	for i, numStr := range numbers {
		// convert string to integer
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Error converting %s to int: %v\n", numStr, err)
			continue
		}

		// alternate placement into arrays
		if i%2 == 0 {
			arr1 = append(arr1, num)
		} else {
			arr2 = append(arr2, num)
		}
	}

	return arr1, arr2
}
