from typing import List


class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        total = sum(nums)
        if total % 2 != 0:
            return False
        subsum = total // 2
        n = len(nums)
        dp = [[False for _ in range(subsum + 1)] for _ in range(n + 1)]
        dp[0][0] = True
        for i in range(1, n + 1):
            cur = nums[i - 1]
            for j in range(0, subsum + 1):
                if j < cur:
                    dp[i][j] = dp[i - 1][j]
                else:
                    dp[i][j] = dp[i - 1][j] or dp[i - 1][j - cur]
        return dp[n][subsum]


if __name__ == "__main__":
    nums = [1, 5, 11, 5]
    # nums = [1, 3, 5]
    nums = [1, 3, 4, 4]
    # nums = [14, 9, 8, 4, 3, 2]
    print(Solution().canPartition(nums))
