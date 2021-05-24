from typing import List


class Solution:
    def minSubArrayLen(self, s: int, nums: List[int]) -> int:
        res = float("inf")
        left = 0
        cur = 0
        for i in range(len(nums)):
            cur += nums[i]
            while cur >= s:
                res = min(res, i + 1 - left)
                cur -= nums[left]
                left += 1
        return res if res < float("inf") else 0


if __name__ == "__main__":
    # s = 11
    # nums = [1, 2, 3, 4, 5]
    s = 11
    nums = [1, 1, 1, 1, 1, 1, 1, 1]
    print(Solution().minSubArrayLen(s, nums))
