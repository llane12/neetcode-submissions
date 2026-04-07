type Node struct {
	val int
	prev *Node
}

func evalRPN(tokens []string) int {
	var stack *Node
	for _, token := range tokens {
		if i, err := strconv.Atoi(token); err == nil {
			stack = &Node{
				val: i,
				prev: stack,
			}
		} else {
			op2 := stack.val
			op1 := stack.prev.val
			var res int
			if token == "+" {
				res = op1 + op2
			} else if token == "-" {
				res = op1 - op2
			} else if token == "*" {
				res = op1 * op2
			} else {
				res = op1 / op2
			}
			stack = &Node{
				val: res,
				prev: stack.prev.prev,
			}
		}
	}
	return stack.val
}
