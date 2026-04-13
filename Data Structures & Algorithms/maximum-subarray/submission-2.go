func maxSubArray(nums []int) int {
    max_sum, cur_sum := nums[0], 0
    for _, num := range nums {
		if cur_sum < 0 {
			cur_sum = 0
		}        
        cur_sum += num
		max_sum = max(max_sum, cur_sum)
    }
    return max_sum
}
