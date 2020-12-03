from typing import List


class Solution:
    def combinationSum2(self, candidates: List[int],
                        target: int) -> List[List[int]]:
        candidates.sort()
        res = []

        def helper(cur, target, path):
            if target == 0:
                res.append(path[:])
                return
            if target < 0:
                return
            for i in range(cur, len(candidates)):
                if i > cur and candidates[i] == candidates[i - 1]:
                    continue
                path.append(candidates[i])
                helper(i + 1, target - candidates[i], path)
                path.pop()

        helper(0, target, [])
        return res
