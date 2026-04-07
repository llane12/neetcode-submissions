func longestConsecutive(nums []int) int {
	// if len(nums) == 0 {
	// 	return 0
	// }

	hm := map[int]bool{}	
	for _, num := range nums {
		hm[num] = true
	}

	var longest int
	for _, num := range nums {
		// ONLY if this number is the start of a new sequence,
		// which reduces from O(n^2) --> O(2n)
		if !hm[num - 1] {
			length := 1
			for hm[num + length] { // no 'while' in Go
				length++
			}
			longest = max(longest, length)
		}
	}
	return longest
}
