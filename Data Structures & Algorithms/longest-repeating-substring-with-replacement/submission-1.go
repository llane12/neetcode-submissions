func characterReplacement(s string, k int) int {
	l, maxLen, maxRepeatCount := 0, 0, 0
	charCounts := make([]int, 26)
	for r, char := range s {
		charIndex := char-'A'
		charCounts[charIndex]++
		maxRepeatCount = max(maxRepeatCount, charCounts[charIndex])

		if (r - l + 1 - maxRepeatCount) > k {
			charCounts[s[l]-'A']--
			l++
		}

		maxLen = max(maxLen, r - l + 1)
	}
	return maxLen
}
