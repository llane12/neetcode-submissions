func twoSum(nums []int, target int) []int {
    positions := map[int]int{}
	for i, n := range nums {
		_, seen1 := positions[n]
		if !seen1 {
			positions[n] = i
		}

		idx, seen2 := positions[target - n]
		if seen2 && idx != i {
			return []int{idx, i}
		}
	}
	return []int{}
}
