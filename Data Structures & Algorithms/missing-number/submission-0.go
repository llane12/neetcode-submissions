func missingNumber(nums []int) int {
    n := len(nums)
    
    xor := 0
    for i := 0; i <= n; i++ {
        xor ^= i
    }

    for _, num := range nums {
        xor ^= num
    }

    return xor
}
