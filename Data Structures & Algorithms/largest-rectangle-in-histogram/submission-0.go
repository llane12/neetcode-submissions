func largestRectangleArea(heights []int) int {
	var max int
	for i := 0; i < len(heights); i++ {		
		//for height := heights[i]; height > 0; height-- {
			height:= heights[i]
			width := 1
			for l := i - 1; l >= 0; l-- {
				if heights[l] < height { break }
				width += 1
			}
			for r := i + 1; r < len(heights); r++ {
				if heights[r] < height { break }
				width += 1
			}
			area := width * height
			if area > max { max =area }
		//}
	}
	return max
}
