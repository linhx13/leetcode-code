from typing import List


class Solution:
    def subarraysDivByK(self, A: List[int], K: int) -> int:
        res = 0
        cnts = {0: 1}
        s = 0
        for n in A:
            s += n
            r = s % K
            res += cnts.get(r, 0)
            cnts[r] = cnts.get(r, 0) + 1
        return res
