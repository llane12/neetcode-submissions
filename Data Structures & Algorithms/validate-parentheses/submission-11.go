type Node struct {
	val rune
	prev *Node
}

func isValid(s string) bool {
	if len(s) < 2 || len(s) % 2 != 0 {
		return false
	}

	pairs := map[rune]rune {
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack *Node
	for _, c := range s {
		switch c {
			case '(', '{', '[':
				stack = &Node{
				val: c,
				prev: stack,
			}
			case ')', '}', ']':
				if stack == nil || stack.val != pairs[c] {
					return false
				}
				stack = stack.prev
			default:
				return false
		}
	}
	return stack == nil
}
