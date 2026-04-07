func trap(height []int) int {
	res := 0
	l, r := 0, len(height) - 1
	maxL, maxR := height[l], height[r]
	for l < r {
		if height[l] < height[r] {
			if height[l] < maxL {
				res += maxL - height[l]
			} else {
				maxL = height[l]
			}
			l++
		} else {
			if height[r] < maxR {
				res += maxR - height[r]
			} else {
				maxR = height[r]
			}
			r--
		}
	}
	return res
}
