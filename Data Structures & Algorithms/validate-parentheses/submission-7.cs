public class Solution {
    public bool IsValid(string s) {
        Stack<char> stack = new Stack<char>();        
        for (int i = 0; i < s.Length; i++)
        {
            if (s[i] == '(' || s[i] == '{' || s[i] == '[')
            {
                stack.Push(s[i]);
            }
            else if (stack.Count > 0)
            {
                char prev = stack.Pop();
                switch ((prev, s[i]))
                {
                    case ('(', ')'):
                    case ('{', '}'):
                    case ('[', ']'):
                        continue;
                    default:
                        return false;
                }
            }
            else
            {
                return false;
            }
        }
        return stack.Count == 0;
    }
}
