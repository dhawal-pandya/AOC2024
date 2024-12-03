package main

import (
	"fmt"
)

func main() {

	result := CalculateSum(Input)
	fmt.Printf("Sum of valid multiplications: %d\n", result) // Sum of valid multiplications: 187194524

	result = CalculateSumWithConditions(Input)
	fmt.Printf("Sum of valid and enabled multiplications: %d\n", result) // Sum of valid and enabled multiplications: 127092535
}
