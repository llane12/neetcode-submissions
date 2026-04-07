class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        results = []
        self.recursive(nums, [], results)        
        return results

    def recursive(self,
        nums: List[int],
        perm: List[int],
        results: List[List[int]]):

        # Base case
        if len(nums) == 0:
            results.append(perm.copy())
            return

        for i, num in enumerate(nums):
            perm.append(num)

            other_nums = nums[:i] + nums[i+1:]
            self.recursive(other_nums, perm, results)

            perm.pop()