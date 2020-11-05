from typing import List


class Solution:
    def maxSumDivThree(self, nums: List[int]) -> int:
        dp = [[0 for _ in range(3)] for _ in range(len(nums))]
        for i in range(len(nums)):
            if i == 0:
                dp[i][nums[i] % 3] = nums[i]
            else:
                for j in range(3):
                    dp[i][j] = dp[i - 1][j]
                for j in range(3):
                    x = nums[i] + dp[i - 1][j]
                    r = x % 3
                    dp[i][r] = max(dp[i][r], x)
        return dp[-1][0]


if __name__ == "__main__":
    # nums = [4]
    nums = [1, 2, 3, 4, 4]
    print(Solution().maxSumDivThree(nums))
