package main

import (
	"bufio"
	"strconv"
	"strings"
)

// CountSafeReports takes input data and counts the number of safe reports
func CountSafeReports(data string) int {
	scanner := bufio.NewScanner(strings.NewReader(data))
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		levels := make([]int, len(parts))
		for i, part := range parts {
			levels[i], _ = strconv.Atoi(part)
		}

		if IsReportSafe(levels) {
			safeCount++
		}
	}

	return safeCount
}
