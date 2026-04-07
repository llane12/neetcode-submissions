func largestRectangleArea(heights []int) int {
	if len(heights) == 0 { return 0 }
	if len(heights) == 1 { return heights[0] }

	type node struct {
		idx, val int
		prev *node
	}

	max := heights[0]
	stack := &node{
		idx: 0,
		val: heights[0],
	}
	for i := 1; i < len(heights); i++ {		
		height:= heights[i]
		idx := i
		for stack != nil && height < stack.val {
			area := (i - stack.idx) * stack.val
			if area > max { max = area }
			idx = stack.idx
			stack = stack.prev
		}
		stack = &node{
			idx: idx,
			val: height,
			prev: stack,
		}
	}
	for stack != nil {
		area := (len(heights) - stack.idx) * stack.val
		if area > max { max = area }
		stack = stack.prev
	}
	return max
}

func largestRectangleArea_BruteForce(heights []int) int {
	var max int
	for i := 0; i < len(heights); i++ {		
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
	}
	return max
}
