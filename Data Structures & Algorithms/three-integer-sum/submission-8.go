import "slices"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)
	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		target := -nums[i]
		for j, k := i + 1, len(nums) - 1; j < k; {
			sum := nums[j] + nums[k]
			if sum < target {
				j++
			} else if sum > target {
				k--
			} else {
				res = append(res, []int{nums[i],nums[j],nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return res
}
