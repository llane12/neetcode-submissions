func lengthOfLongestSubstring(s string) int {
	//if len(s) == 0 {return 0}
	//if len(s) == 1 {return 1}
	//runes := []rune(s)

	substr := make(map[rune]bool)
	maxLen, l := 0, 0
	for r, char := range s {
		if _, found := substr[char]; found {
			for substr[char] {
				delete(substr, rune(s[l]))
				l++
			}
		}
		substr[char] = true
		maxLen = max(maxLen, r - l + 1)
	}
	return maxLen
}
