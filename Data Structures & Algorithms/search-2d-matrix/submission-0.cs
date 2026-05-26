public class Solution {
    public bool SearchMatrix(int[][] matrix, int target) {
        int t = 0, b = matrix.Length - 1;
        int row = 0;
        while (t <= b)
        {
            int mid = t + ((b - t) / 2);
            if (matrix[mid][0] == target)
            {
                return true;
            }
            else if (matrix[mid][0] < target)
            {
                row = mid;
                t = mid + 1;
            }
            else
            {
                b = mid - 1;
            }
        }

        int l = 0, r = matrix[row].Length - 1;
        while (l <= r)
        {
            int mid = l + ((r - l) / 2);
            if (matrix[row][mid] == target)
            {
                return true;
            }
            else if (matrix[row][mid] < target)
            {
                l = mid + 1;
            }
            else
            {
                r = mid - 1;
            }
        }

        return false;
    }
}
