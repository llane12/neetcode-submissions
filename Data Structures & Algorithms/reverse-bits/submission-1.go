func reverseBits(n int) int {
	// Each time, get the least significant bit from the input and add it to the result
	// then shift the result up by 1, so by the end the first bit becomes the last.
    res := 0
	ncopy := n
    for range 32 {
        res <<= 1
        res |= (ncopy & 1)
        ncopy >>= 1
    }    
    return res
}
