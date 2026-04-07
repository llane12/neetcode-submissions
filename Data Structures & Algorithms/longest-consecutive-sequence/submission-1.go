func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := map[int]bool{}	
	for _, n := range nums {
		count[n] = true
	}

	longest := 0
	for _, num := range nums {
		// ONLY if this number is the start of a new sequence,
		// which reduces from O(n^2) --> O(2n)
		if !count[num - 1] {
			length := 1
			for count[num + length] { // no 'while' in Go
				length++
			}
			longest = max(longest, length)
		}
	}
	return longest
}
