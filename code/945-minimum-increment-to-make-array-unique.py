from typing import List


class Solution:
    def minIncrementForUnique(self, A: List[int]) -> int:
        A.sort()
        target = 0
        ans = 0
        for n in A:
            if n >= target:
                target = n + 1
            else:
                ans += target - n
                target += 1
        return ans
