from typing import List


class Solution:
    def maxSumAfterPartitioning(self, arr: List[int], k: int) -> int:
        dp = [0] * (len(arr) + 1)
        for i in range(1, len(arr) + 1):
            for j in range(1, k + 1):
                if i - j < 0:
                    break
                dp[i] = max(dp[i], dp[i - j] + max(arr[i - j:i]) * j)
        return dp[len(arr)]


if __name__ == "__main__":
    arr = [1, 15, 7, 9, 2, 5, 10]
    # arr = [1, 15, 7]
    k = 3
    sol = Solution()
    print(sol.maxSumAfterPartitioning(arr, k))
