type Node struct {
	prev *Node
	idx int
	val int
}

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	} else if len(temperatures) == 1 {
		return []int{0}
	}

	res := make([]int, len(temperatures))
	stack := &Node{
		idx: 0,
		val: temperatures[0],
	}
	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] > stack.val {
			for stack != nil && temperatures[i] > stack.val {
				res[stack.idx] = i - stack.idx
				stack = stack.prev
			}
		}
		stack = &Node{
			idx: i,
			val: temperatures[i],
			prev: stack,
		}
	}
	return res
}
