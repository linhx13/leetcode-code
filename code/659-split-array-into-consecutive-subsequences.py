from typing import List
import collections


class Solution:
    def isPossible(self, nums: List[int]) -> bool:
        counter = collections.Counter(nums)
        tails = collections.Counter()
        for x in nums:
            if counter[x] == 0:
                continue
            elif tails[x] > 0:
                tails[x] -= 1
                tails[x + 1] += 1
            elif counter[x + 1] > 0 and counter[x + 2] > 0:
                counter[x + 1] -= 1
                counter[x + 2] -= 1
                tails[x + 3] += 1
            else:
                return False
            counter[x] -= 1
        return True
