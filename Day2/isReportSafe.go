package main

// IsReportSafe checks if a report is safe based on the given rules
func IsReportSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	increasing := levels[1] > levels[0]
	decreasing := levels[1] < levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		// check difference is within the allowed range and not zero
		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}

		// check consistency in increasing or decreasing trend
		if diff > 0 {
			if decreasing {
				return false
			}
			increasing = true
		} else if diff < 0 {
			if increasing {
				return false
			}
			decreasing = true
		}
	}

	return true
}
