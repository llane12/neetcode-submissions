func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := map[int]int{}
	start := 0
	end := 0
	
	for i, n := range nums {
		if i == 0 {
			start = n
			end = n
		} else {
			if n < start {
				start = n
			} else if n > end {
				end = n
			}
		}
		count[n]++
	}

	longest := 0
	length := 0

	for n := start; n <= end; n++ {
		_, found := count[n]
		if found {
			length++
			if length > longest {
				longest = length
			}
		} else {
			length = 0
		}
	}

	return longest
}
