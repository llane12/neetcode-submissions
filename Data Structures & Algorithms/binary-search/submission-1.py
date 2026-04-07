class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums) - 1
        while True:
            m = int((l + r) / 2)
            if nums[m] == target:
                return m
            elif nums[r] == target:
                return r
            elif l == m:
                return -1
            elif nums[m] < target:
                l = m
            else:
                r = m