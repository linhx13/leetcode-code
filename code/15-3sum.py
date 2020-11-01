from typing import List
from collections import defaultdict, Counter


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        if len(nums) < 3:
            return {}
        nums.sort()
        res = []
        for i in range(len(nums) - 3 + 1):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            l = i + 1
            r = len(nums) - 1
            s = -nums[i]
            cur = nums[l] + nums[r]
            while l < r:
                if cur == s:
                    res.append([nums[i], nums[l], nums[r]])
                if cur <= s:
                    cur -= nums[l]
                    l += 1
                    while l < r and nums[l] == nums[l - 1]:
                        l += 1
                    if l < r:
                        cur += nums[l]
                elif cur > s:
                    cur -= nums[r]
                    r -= 1
                    while r > l and nums[r] == nums[r + 1]:
                        r -= 1
                    if r > l:
                        cur += nums[r]
        return res


if __name__ == '__main__':
    nums = [-1, 0, 1, 2, -1, -4]
    # nums = [0]
    print(Solution().threeSum(nums))
