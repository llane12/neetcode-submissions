public class Solution {
    public string MinWindow(string s, string t) {
		int k = t.Length;

        int[] tMap = new int[58];
		for (int i = 0; i < k; i++)
		{
			tMap[t[i] - 'A']++;
		}

		int minLen = int.MaxValue;
		string sub = "";

		int[] sMap = new int[58];
		int l = 0;
		for (int r = 0; r < s.Length; r++)
		{
			sMap[s[r] - 'A']++;

			while (Substring(sMap, tMap))
			{
				int len = r - l + 1;
				if (len < minLen)
				{
					minLen = len;
					sub = s.Substring(l, len);
				}

				sMap[s[l] - 'A']--;
				l++;
			}
		}

		return sub;
    }

	private bool Substring(int[] sMap, int[] tMap)
	{
		for (int i = 0; i < sMap.Length; i++)
		{
			if (sMap[i] < tMap[i]) return false;
		}

		return true;
	}
}
