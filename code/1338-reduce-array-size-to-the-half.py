from typing import List
from collections import Counter


class Solution:
    def minSetSize(self, arr: List[int]) -> int:
        counter = Counter(arr)
        res = 0
        cur = 0
        for k, v in counter.most_common():
            cur += v
            res += 1
            if cur * 2 >= len(arr):
                break
        return res


if __name__ == "__main__":
    # arr = [3, 3, 3, 3, 5, 5, 5, 2, 2, 7]
    # arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    arr = [1000, 1000, 3, 7]
    sol = Solution()
    print(sol.minSetSize(arr))
