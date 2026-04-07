func lengthOfLongestSubstring_map(s string) int {
	substr := make(map[rune]bool, 95) // 95 printable ASCII characters
	maxLen, l := 0, 0
	for r, char := range s {
		if _, found := substr[char]; found {
			for substr[char] { // while
				delete(substr, rune(s[l])) // All ASCII characters are single byte
				l++
			}
		}
		substr[char] = true
		maxLen = max(maxLen, r - l + 1)
	}
	return maxLen
}

func lengthOfLongestSubstring(s string) int {
	// lastSeen stores the index + 1 of each character
	lastSeen := make([]int, 128) 
	maxLen, l := 0, 0

	for r, char := range s {
		// If we've seen this char before AND it's inside our current window
		if lastSeen[char] > l {
			l = lastSeen[char] // Jump 'l' to the position after the duplicate
		}

		lastSeen[char] = r + 1 // Store next possible starting position
		maxLen = max(maxLen, r - l + 1)
	}
	return maxLen
}
