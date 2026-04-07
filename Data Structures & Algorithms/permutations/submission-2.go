func permute(nums []int) [][]int {
    var results [][]int
	find_permutations(nums, []int{}, &results)
	return results
}

func find_permutations(nums, perm []int, results *[][]int) {
	if len(nums) == 0 {
		// Append a copy of the current permutation to the results.
		// It's crucial to copy the slice, as 'perm' is modified in place throughout the recursion.
		*results = append(*results, append([]int{}, perm...))
		return
	}

	for i, num := range nums {
		// 1. Choose: Add the current number to the permutation slice.
		perm = append(perm, num)

		// 2. Create a new slice of remaining numbers (excluding the current one).
		otherNums := make([]int, 0, len(nums)-1)
		otherNums = append(otherNums, nums[:i]...)
		otherNums = append(otherNums, nums[i+1:]...)

		// 3. Recursively call the function with the remaining numbers and updated permutation.
		find_permutations(otherNums, perm, results)

		// 4. Backtrack: remove the last added number from the permutation slice.
		perm = perm[:len(perm)-1]
	}
}
