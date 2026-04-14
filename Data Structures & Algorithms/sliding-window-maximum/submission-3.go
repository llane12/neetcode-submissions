func maxSlidingWindow(nums []int, k int) []int {
    // The deque holds array indexes but is ordered by values
    // (we just can't see them in the slice).
    // We make sure the values are always decreasing,
    // and because we add values left-to-right,
    // the index values will only ever be increasing.
    // So if the left value is in the deque, it will always be at the Front().
    res := make([]int, len(nums) - k + 1)

    dq := list.New()

    // Pre-fill the first window
    for i := 0; i < k; i++ {
        // Remove elements from the back which are smaller than the next
        for dq.Len() > 0 && nums[dq.Back().Value.(int)] < nums[i] {
            dq.Remove(dq.Back())
        }
        dq.PushBack(i)
    }

    for l := 0; l <= len(nums) - k; l++ {
        r := l + k - 1;

        // Record max for this window
        res[l] = nums[dq.Front().Value.(int)]

        // Remove l if it's the front (it's leaving the window)
        if l == dq.Front().Value.(int) {
            dq.Remove(dq.Front())
        }

        // Add the next element if there is one
        if r + 1 < len(nums) {
            // Remove elements from the back which are smaller than the next
            for dq.Len() > 0 && nums[dq.Back().Value.(int)] < nums[r + 1] {
                dq.Remove(dq.Back())
            }
            dq.PushBack(r + 1)
        }
    }

    return res;
}
