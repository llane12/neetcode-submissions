func reverseBits(n int) int {
    res := 0
	ncopy := n
    for range 32 {
        res <<= 1
        res |= (ncopy & 1)
        ncopy >>= 1
    }    
    return res
}
