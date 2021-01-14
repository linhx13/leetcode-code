from typing import List


class Solution:
    def numTimesAllBlue(self, light: List[int]) -> int:
        right = res = 0
        for i, x in enumerate(light, 1):
            right = max(right, x)
            if right == i:
                res += 1
        return res


if __name__ == "__main__":
    # light = [2, 1, 3, 5, 4]
    # light = [3, 2, 4, 1, 5]
    # light = [4, 1, 2, 3]
    # light = [2, 1, 4, 3, 6, 5]
    light = [1, 2, 3, 4, 5, 6]
    # light = [1]
    print(Solution().numTimesAllBlue(light))
