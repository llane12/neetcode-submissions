class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if t == "" or len(s) < len(t): return ""

        countT, window = {}, {}
        for c in t:
            countT[c] = 1 + countT.get(c, 0)

        have, need = 0, len(countT)
        res, resLen = [-1, -1], float("infinity")

        l = 0
        for r in range(len(s)):
            c = s[r]
            window[c] = 1 + window.get(c, 0)
            
            if c in countT and window[c] == countT[c]:
                have += 1
            
            while have == need:
                if (r - l + 1) < resLen:
                    res = [l, r]
                    resLen = (r - l + 1)
                # pop from left of window
                cl = s[l]
                window[cl] -= 1 # update window map
                if cl in countT and window[cl] < countT[cl]:
                    have -= 1 # we just lost this character
                l += 1
        
        l, r = res
        if resLen != float("infinity"):
            return s[l:r + 1]
        else:
            return ""
