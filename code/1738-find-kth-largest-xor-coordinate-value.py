from typing import List


class Solution:
    def kthLargestValue(self, matrix: List[List[int]], k: int) -> int:
        m, n = len(matrix), len(matrix[0])
        sums = [0] * (m * n)
        for i in range(m):
            for j in range(n):
                s = matrix[i][j]
                if i - 1 >= 0:
                    s ^= sums[n * (i - 1) + j]
                if j - 1 >= 0:
                    s ^= sums[n * i + j - 1]
                if i - 1 >= 0 and j - 1 >= 0:
                    s ^= sums[n * (i - 1) + j - 1]
                sums[n * i + j] = s
        sums.sort()
        return sums[-k]


if __name__ == "__main__":
    matrix = [[5, 2], [1, 6]]
    k = 2
    print(Solution().kthLargestValue(matrix, k))
