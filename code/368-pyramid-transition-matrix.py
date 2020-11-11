from typing import List


class Solution:
    def largestDivisibleSubset(self, nums: List[int]) -> List[int]:
        if not nums:
            return []
        nums.sort()
        dp = [(1, -1) for i in range(len(nums))]
        max_idx = 0
        for i in range(1, len(nums)):
            for j in range(0, i):
                if nums[i] % nums[j] == 0:
                    if dp[j][0] + 1 > dp[i][0]:
                        dp[i] = (dp[j][0] + 1, j)
            if dp[i][0] > dp[max_idx][0]:
                max_idx = i
        res = []
        while max_idx != -1:
            res.append(nums[max_idx])
            max_idx = dp[max_idx][1]
        res.reverse()
        return res


if __name__ == "__main__":
    # nums = [2, 3, 6, 8, 12]
    nums = []
    print(Solution().largestDivisibleSubset(nums))
