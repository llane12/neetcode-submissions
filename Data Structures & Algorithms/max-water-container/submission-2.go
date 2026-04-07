func maxArea(heights []int) int {
	area := 0

	for i, j := 0, len(heights) - 1; i < j; {
		newArea := (j - i) * min(heights[i], heights[j])
		area = max(area, newArea)

		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return area
}
