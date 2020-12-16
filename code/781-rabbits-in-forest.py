from typing import List
from collections import Counter


class Solution:
    def numRabbits(self, answers: List[int]) -> int:
        counter = Counter(answers)
        res = 0
        for k, v in counter.items():
            group_size = k + 1
            res += (v + group_size - 1) // group_size * group_size
        return res
