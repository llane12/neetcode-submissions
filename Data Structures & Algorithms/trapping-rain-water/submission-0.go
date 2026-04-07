func trap(height []int) int {
	m := 0
	maxLeft := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		m = max(m, height[i])
		maxLeft[i] = m
	}

	m = 0
	maxRight := make([]int, len(height))
	for i := len(height) - 1; i >= 0; i-- {
		m = max(m, height[i])
		maxRight[i] = m
	}

	area := 0
	for i := 0; i < len(height); i++ {
		a := min(maxLeft[i], maxRight[i]) - height[i]
		if a > 0 {
			area += a
		}
	}
	return area
}
