func trap(height []int) int {
	res := 0

	m := 0
	maxL := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		m = max(m, height[i])
		maxL[i] = m
	}

	m = 0
	maxR := make([]int, len(height))
	for i := len(height) - 1; i >= 0; i-- {
		m = max(m, height[i])
		maxR[i] = m
	}

	for i := 0; i < len(height); i++ {
		a := min(maxL[i], maxR[i]) - height[i]
		if a > 0 {
			res += a
		}
	}

	return res
}
