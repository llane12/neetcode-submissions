func lengthOfLongestSubstring(s string) int {
	substr := make(map[rune]bool, 95) // 95 printable ASCII characters
	maxLen, l := 0, 0
	for r, char := range s {
		if _, found := substr[char]; found {
			for substr[char] {
				delete(substr, rune(s[l])) // All ASCII characters are single byte
				l++
			}
		}
		substr[char] = true
		maxLen = max(maxLen, r - l + 1)
	}
	return maxLen
}
