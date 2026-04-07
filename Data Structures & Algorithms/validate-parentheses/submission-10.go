type Node struct {
	val rune
	prev *Node
}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	var stack *Node
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = &Node{
				val: char,
				prev: stack,
			}
		} else if stack == nil {
			return false
		} else {
			switch char {
				case ')':
					if stack.val != '(' {
						return false
					}
				case '}':
					if stack.val != '{' {
						return false
					}
				case ']':
					if stack.val != '[' {
						return false
					}
			}
			stack = stack.prev
		}	
	}
	return stack == nil
}
