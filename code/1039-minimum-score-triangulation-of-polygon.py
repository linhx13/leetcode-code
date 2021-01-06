from typing import List


class Solution:
    def minScoreTriangulation(self, A: List[int]) -> int:
        n = len(A)
        dp = [[0 for _ in range(n)] for _ in range(n)]
        for d in range(2, n):
            for i in range(n - d):
                j = i + d
                dp[i][j] = min(
                    dp[i][k] + dp[k][j] + A[i] * A[j] * A[k]
                    for k in range(i + 1, j)
                )
        return dp[0][n - 1]


if __name__ == "__main__":
    # A = [1, 2, 3]
    # A = [3, 7, 4, 5]
    # A = [1, 3, 1, 4, 1, 5]
    # A = [2, 1, 4, 4]
    A = [
        20,
        3,
        42,
        70,
        54,
        40,
        54,
        65,
        48,
        93,
        86,
        100,
        75,
        100,
        24,
        3,
        46,
        54,
        82,
        11,
        94,
        33,
        75,
        32,
        20,
        35,
        49,
        47,
        46,
        96,
        43,
        76,
        38,
        38,
        51,
        86,
        5,
        19,
        30,
        73,
        66,
        2,
        55,
        23,
        32,
        13,
        86,
        100,
        95,
        24,
    ]

    print(Solution().minScoreTriangulation(A))
