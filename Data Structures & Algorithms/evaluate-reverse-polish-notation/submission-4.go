func evalRPN(tokens []string) int {
    if len(tokens) == 0 {
        return 0
    }

    operators := map[string]func(a, b int) int{
        "+": func(a, b int) int { return a + b },
        "-": func(a, b int) int { return a - b },
        "*": func(a, b int) int { return a * b },
        "/": func(a, b int) int { return a / b },
    }

    stack := make([]int, 0)

    for _, token := range tokens {
        if op, found := operators[token]; found {
            var op1, op2 int
            stack, op2 = pop(stack)
            stack, op1 = pop(stack)
            
            res := op(op1, op2)
            stack = append(stack, res)
        } else {
            num, _ := strconv.Atoi(token)
            stack = append(stack, num)
        }
    }

    return stack[0]
}

func pop(stack []int) ([]int, int) {
    return stack[:len(stack) - 1], stack[len(stack) - 1]
}
