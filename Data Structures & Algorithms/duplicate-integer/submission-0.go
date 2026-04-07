func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    for _, num := range nums {
        _, alreadySeen := seen[num]
        if alreadySeen {
            return true
        }
        seen[num] = true
    }
    return false
}
