from typing import List, Set, Tuple


class Solution:
    def findSubsequences(self, nums: List[int]) -> List[List[int]]:
        res: List[List[int]] = []
        memo: Set[Tuple[int]] = set()
        self.helper(nums, 0, [], res, memo)
        return res

    def helper(self, nums, idx, cur, res, memo):
        if idx >= len(nums):
            return
        last = -200
        if len(cur) != 0:
            last = cur[-1]
        for i in range(idx, len(nums)):
            if nums[i] >= last:
                cur.append(nums[i])
                if len(cur) >= 2 and tuple(cur) not in memo:
                    tmp = [x for x in cur]
                    res.append(tmp)
                    memo.add(tuple(tmp))
                self.helper(nums, i + 1, cur, res, memo)
                cur.pop()
