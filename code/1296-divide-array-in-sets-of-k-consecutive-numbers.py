from typing import List
from collections import Counter


class Solution:
    def isPossibleDivide(self, nums: List[int], k: int) -> bool:
        if len(nums) % k != 0:
            return False
        counter = Counter(nums)
        nums = sorted(counter.keys())
        i = 0
        while i < len(nums):
            if counter[nums[i]] == 0:
                while i < len(nums) and counter[nums[i]] == 0:
                    i += 1
                continue
            for x in range(nums[i] + 1, nums[i] + k):
                if counter[x] == 0:
                    return False
                counter[x] -= 1
            counter[nums[i]] -= 1
        return True
