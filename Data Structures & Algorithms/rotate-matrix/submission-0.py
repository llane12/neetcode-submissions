class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)
        for r in range(n // 2):
            for c in range(n):
                # vertical mirror
                print(f"mirroring [{r},{c}] and [{n-1-r},{c}]")
                t = matrix[n-1-r][c]
                matrix[n-1-r][c] = matrix[r][c]
                matrix[r][c] = t

        for r in range(n):
            for c in range(n):
                # diagonal mirror / transpose
                if c > r:
                    print(f"transposing [{r},{c}] and [{c},{r}]")
                    t = matrix[c][r]
                    matrix[c][r] = matrix[r][c]
                    matrix[r][c] = t