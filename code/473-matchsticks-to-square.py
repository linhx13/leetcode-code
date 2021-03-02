from typing import List
from functools import lru_cache


class Solution:
    def makesquare(self, nums: List[int]) -> bool:
        if len(nums) < 4:
            return False

        nums.sort()
        limit = sum(nums) >> 2
        if limit * 4 != sum(nums):
            return False
        if any(x > limit for x in nums):
            return False

        @lru_cache(None)
        def dfs(i, sides):
            if i == -1:
                return (
                    0 != sides[0]
                    and sides[0] == sides[1]
                    and sides[1] == sides[2]
                    and sides[2] == sides[3]
                )
            sides = list(sides)
            for j in range(len(sides)):
                if sides[j] + nums[i] <= limit:
                    sides[j] += nums[i]
                    if dfs(i - 1, tuple(sorted(sides))):
                        return True
                    sides[j] -= nums[i]
            return False

        return dfs(len(nums) - 1, (0, 0, 0, 0))


if __name__ == "__main__":
    # nums = [7, 4, 2, 18, 13, 6, 10, 3, 8, 4, 11, 14, 6, 13, 11]
    # nums = [
    #     2797277,
    #     8645366,
    #     329762,
    #     8136656,
    #     4696913,
    #     8549690,
    #     1480677,
    #     9584058,
    #     6064920,
    #     8155913,
    #     8012446,
    #     2032022,
    #     4943218,
    #     2849467,
    #     91762,
    # ]
    nums = [
        1569462,
        2402351,
        9513693,
        2220521,
        7730020,
        7930469,
        1040519,
        5767807,
        876240,
        350944,
        4674663,
        4809943,
        8379742,
        3517287,
        8034755,
    ]
    print(Solution().makesquare(nums))
