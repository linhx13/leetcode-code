from typing import List
import functools


class Solution:
    def canCross(self, stones: List[int]) -> bool:
        @functools.lru_cache(None)
        def dfs(p, k):
            if p == len(stones) - 1:
                return True
            for i in range(p + 1, len(stones)):
                gap = stones[i] - stones[p]
                if gap < k - 1:
                    continue
                if gap > k + 1:
                    return False
                if dfs(i, gap):
                    return True
            return False

        return dfs(0, 0)


if __name__ == "__main__":
    stones = [0, 2]
    print(Solution().canCross(stones))
