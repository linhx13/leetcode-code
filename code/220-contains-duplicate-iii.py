from typing import List
from collections import OrderedDict


class Solution:
    def containsNearbyAlmostDuplicate(self, nums: List[int], k: int,
                                      t: int) -> bool:
        od = OrderedDict()
        window = t + 1
        for i, n in enumerate(nums):
            while len(od) > k:
                od.popitem(last=False)
            b = n // window
            if b in od:
                return True
            if b - 1 in od and abs(od[b - 1] - n) <= t:
                return True
            if b + 1 in od and abs(od[b + 1] - n) <= t:
                return True
            od[b] = n
        return False
