from typing import List
import collections


class Solution:
    def countPairs(self, deliciousness: List[int]) -> int:
        freq = collections.Counter(deliciousness)
        res = 0
        for x in freq:
            for k in range(22):
                t = 1 << k
                if t - x <= x and t - x in freq:
                    if t - x == x:
                        res += freq[x] * (freq[x] - 1) // 2
                    else:
                        res += freq[x] * freq[t - x]
        return res % (10 ** 9 + 7)
