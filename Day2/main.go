package main

import (
	"fmt"
)

func main() {

	safeReports := CountSafeReports(Input)
	fmt.Println("Number of Safe Reports:", safeReports) // Number of Safe Reports: 246

	safeReports = CountSafeReportsWithDampener(Input)
	fmt.Println("Number of Safe Reports with Dampener:", safeReports) // Number of Safe Reports with Dampener: 318
}
