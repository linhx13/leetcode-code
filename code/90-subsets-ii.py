from typing import List


class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        res = set()
        N = len(nums)

        def dfs(i, cur):
            if i == N:
                res.add(tuple(sorted(cur)))
                return
            dfs(i + 1, cur)
            cur.append(nums[i])
            dfs(i + 1, cur)
            cur.pop()

        dfs(0, [])

        return res
