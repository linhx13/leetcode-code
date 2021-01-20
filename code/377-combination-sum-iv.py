from typing import List
import functools


class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        nums.sort()

        @functools.lru_cache(None)
        def dfs(target):
            res = 0
            for x in nums:
                if x == target:
                    res += 1
                elif x > target:
                    break
                else:
                    res += dfs(target - x)
            return res

        return dfs(target)


if __name__ == "__main__":
    nums = [1, 2, 3]
    target = 4
    print(Solution().combinationSum4(nums, target))
