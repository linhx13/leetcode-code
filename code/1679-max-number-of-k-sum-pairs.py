from typing import List
import collections


class Solution:
    def maxOperations(self, nums: List[int], k: int) -> int:
        counter = collections.Counter(nums)
        res = 0
        for x, c in counter.items():
            if (x << 1) != k and (k - x) in counter:
                res += min(c, counter[k - x])
        res >>= 1
        if (k & 1 == 0) and (k >> 1) in counter:
            res += counter[k >> 1] >> 1
        return res


if __name__ == "__main__":
    nums = [3, 1, 3, 4, 3]
    k = 6
    print(Solution().maxOperations(nums, k))
