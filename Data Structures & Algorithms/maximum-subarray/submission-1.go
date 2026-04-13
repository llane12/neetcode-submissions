func maxSubArray(nums []int) int {
    max_sum, cur_sum := nums[0], 0
    for _, num := range nums {
        cur_sum += num
		if cur_sum > max_sum {
			max_sum = cur_sum
		}
		if cur_sum < 0 {
			cur_sum = 0
		}        
    }
    return max_sum
}
