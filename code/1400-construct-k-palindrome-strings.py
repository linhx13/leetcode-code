import collections


class Solution:
    def canConstruct(self, s: str, k: int) -> bool:
        if len(s) < k:
            return False
        counter = collections.Counter(s)
        x = sum(1 for v in counter.values() if v & 1 == 1)
        return x <= k
