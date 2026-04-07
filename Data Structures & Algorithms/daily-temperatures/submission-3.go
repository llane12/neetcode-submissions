type Node struct {
	idx, val int
	prev *Node
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack *Node
	for i, temp := range temperatures {
		for stack != nil && temp > stack.val {
			res[stack.idx] = i - stack.idx
			stack = stack.prev
		}
		stack = &Node{
			idx: i,
			val: temp,
			prev: stack,
		}
	}
	return res
}
