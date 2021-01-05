from typing import List


class Solution:
    def findTargetSumWays(self, nums: List[int], S: int) -> int:
        s = sum(nums)
        if S > s or s < -S:
            return 0
        if (S + s) & 1 == 1:
            return 0
        target = (s + S) >> 1
        dp = [0 for _ in range(target + 1)]
        dp[0] = 1
        for i in range(1, len(nums) + 1):
            for j in range(target, nums[i - 1] - 1, -1):
                dp[j] += dp[j - nums[i - 1]]
        return dp[target]


if __name__ == "__main__":
    nums = [1]
    S = 1
    sol = Solution()
    print(sol.findTargetSumWays(nums, S))
