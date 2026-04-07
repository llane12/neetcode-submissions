class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)

        # vertical mirror
        for r in range(n // 2):
            for c in range(n):
                t = matrix[n-1-r][c]
                matrix[n-1-r][c] = matrix[r][c]
                matrix[r][c] = t

        # diagonal mirror / transpose
        for r in range(n):
            for c in range(r + 1, n):
                    t = matrix[c][r]
                    matrix[c][r] = matrix[r][c]
                    matrix[r][c] = t