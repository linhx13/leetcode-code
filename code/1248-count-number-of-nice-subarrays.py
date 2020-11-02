from typing import List
from collections import Counter


class Solution:
    def numberOfSubarrays(self, nums: List[int], k: int) -> int:
        prefix = [0]
        for i in nums:
            prefix.append(prefix[-1] + i % 2)
        counter = Counter(prefix)
        res = 0
        for i, j in counter.items():
            res += j * counter.get(i + k, 0)
        return res


if __name__ == "__main__":
    # nums = [1, 1, 2, 1, 1]
    # k = 3
    # nums = [2, 4, 6]
    # k = 1
    nums = [2, 2, 2, 1, 2, 2, 1, 2, 2, 2]
    k = 2
    print(Solution().numberOfSubarrays(nums, k))
