public class Solution {
    public bool IsValid(string s) {
        var stack = new Stack<char>();
        foreach (char c in s)
        {
            try
            {
                if (c == '(' || c == '{' || c == '[')
                {
                    stack.Push(c);
                }
                else if (c == ')')
                {
                    char prev = stack.Pop();
                    if (prev != '(')
                        return false;
                }
                else if (c == '}')
                {
                    char prev = stack.Pop();
                    if (prev != '{')
                        return false;
                }
                else if (c == ']')
                {
                    char prev = stack.Pop();
                    if (prev != '[')
                        return false;
                }
                else
                {
                    return false;
                }
            }
            catch (InvalidOperationException)
            {
                // Empty stack - closing bracket with no opens
                return false;
            }
        }
        return stack.Count == 0;
    }
}
