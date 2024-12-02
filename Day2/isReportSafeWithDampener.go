package main

// IsSafeWithDampener checks if a report can be made safe by removing one level
func IsSafeWithDampener(levels []int) bool {
	if IsReportSafe(levels) {
		return true
	}

	// try removing each level and check if the report becomes safe
	for i := range levels {
		modifiedLevels := append([]int{}, levels[:i]...)         // copy levels up to i
		modifiedLevels = append(modifiedLevels, levels[i+1:]...) // append levels after i
		if IsReportSafe(modifiedLevels) {
			return true
		}
	}

	return false
}
