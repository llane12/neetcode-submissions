public class Solution {
    public bool IsValid(string s) {
        var stack = new Stack<char>();
        foreach (char c in s)
        {
            if (c == '(' || c == '{' || c == '[')
            {
                stack.Push(c);
            }
            else if (stack.Count > 0)
            {
                char prev = stack.Pop();
                if (c == ')' && prev != '(') return false;
                if (c == '}' && prev != '{') return false;
                if (c == ']' && prev != '[') return false;
            }
            else
            {
                return false;
            }
        }
        return stack.Count == 0;
    }
}
