type Node struct {
	prev *Node
	idx int
	val int
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack *Node
	// stack := &Node{
	// 	idx: 0,
	// 	val: temperatures[0],
	// }
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
