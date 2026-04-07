class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        if len(tokens) == 0: return 0
        
        operators = ("+", "-", "*", "/")
        stack = []

        for token in tokens:
            if token in operators:
                op2 = stack.pop()
                op1 = stack.pop()
                match token:
                    case "+": res = op1 + op2
                    case "-": res = op1 - op2
                    case "*": res = op1 * op2
                    case "/": res = int(op1 / op2)
                stack.append(res)
            else:
                stack.append(int(token, 10))
        return stack.pop()
