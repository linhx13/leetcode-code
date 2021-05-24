from typing import List
import functools


class Solution:
    def arraySign(self, nums: List[int]) -> int:
        def sign(x):
            if x < 0:
                return -1
            elif x > 0:
                return 1
            else:
                return 0

        return functools.reduce(lambda x, y: x * y, map(sign, nums), 1)
