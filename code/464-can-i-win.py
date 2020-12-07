import functools


class Solution:
    def canIWin(self, maxChoosableInteger: int, desiredTotal: int) -> bool:
        if maxChoosableInteger * (maxChoosableInteger + 1) // 2 < desiredTotal:
            return False

        @functools.lru_cache(None)
        def helper(total, remaining):
            nonlocal desiredTotal
            if (remaining[-1] if remaining else 0) + total >= desiredTotal:
                return True
            return any(not helper(total + remaining[i], remaining[:i] +
                                  remaining[i + 1:])
                       for i in range(len(remaining)))

        return helper(0, tuple(range(1, maxChoosableInteger + 1)))
