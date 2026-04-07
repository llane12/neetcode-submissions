func twoSum(nums []int, target int) []int {
    positions := map[int]int{}
	for i, n := range nums {		
		if j, seen := positions[target - n]; seen {
			return []int{j, i}
		}
		positions[n] = i
	}
	return []int{}
}
