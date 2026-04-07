public class Solution {

    public string Encode(IList<string> strs) {
        string encoded = "";
        foreach (string s in strs)
        {
            encoded += $"{s.Length}#~#{s}";
        }
        Console.WriteLine($"Encoded: {encoded}");
        return encoded;
    }

    public List<string> Decode(string s) {
        var strs = new List<string>();
        if (s.Length > 0)
        {
            int i = 0;
            string temp = "";
            while (i < s.Length)
            {
                if (i <= s.Length - 3 && s[i] == '#' && s[i + 1] == '~' && s[i + 2] == '#' && int.TryParse(temp, out int count))
                {
                    strs.Add(s.Substring(i + 3, count));
                    temp = "";
                    i += 3 + count;
                }
                else
                {
                    temp += s[i];
                    i++;
                }
            }
        }
        return strs;
   }
}
