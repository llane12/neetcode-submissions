class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        results = []
        nums.sort()
        for i in range(len(nums) - 2):
            if nums[i] > 0:
                break
            if i > 0 and nums[i] == nums[i - 1]:
                continue

            target = -nums[i]
            j, k = i + 1, len(nums) - 1
            while j < k:
                cur_sum = nums[j] + nums[k]
                if cur_sum < target:
                    j += 1
                elif cur_sum > target:
                    k -= 1
                else:
                    candidate = [nums[i], nums[j], nums[k]]
                    if len(results) == 0 or results[len(results) - 1] != candidate:
                        results.append(candidate)
                    j += 1
        return results