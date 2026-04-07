public class Solution {
    public int[] ProductExceptSelf(int[] nums) {
        int[] output = new int[nums.Length];
        for (int i = 0; i < nums.Length; i++)
        {
            int? product = null;
            for (int j = 0; j < nums.Length; j++)
            {
                if (i == j) continue;
                if (product == null)
                    product = nums[j];
                else
                    product *= nums[j];
            }
            output[i] = product.Value;
        }
        return output;
    }
}
