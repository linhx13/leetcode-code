from typing import List


class Solution:
    def arrayNesting(self, nums: List[int]) -> int:
        visited = [False] * len(nums)
        res = 0
        for i in range(len(nums)):
            if not visited[i]:
                x = nums[i]
                cnt = 0
                while True:
                    cnt += 1
                    visited[x] = True
                    x = nums[x]
                    if x == nums[i]:
                        break
            res = max(res, cnt)
        return res
