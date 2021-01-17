import random


class Solution:
    def __init__(self, nums: List[int]):
        self.nums = nums

    def pick(self, target: int) -> int:
        idx = 0
        cnt = 0
        for i in range(len(self.nums)):
            if self.nums[i] == target:
                cnt += 1
                if random.randint(0, cnt - 1) == 0:
                    idx = i
        return idx
