from typing import List
from collections import Counter


class Solution:
    def canArrange(self, arr: List[int], k: int) -> bool:
        counter = Counter()
        for x in arr:
            counter[((x % k) + k) % k] += 1
        for i in range(k):
            if i == 0 or i == k - i:
                if counter[i] % 2 != 0:
                    return False
            else:
                if counter[i] != counter[k - i]:
                    return False
        return True


if __name__ == "__main__":
    # arr = [1, 2, 3, 4, 5, 10, 6, 7, 8, 9]
    # k = 5
    # arr = [1, 2, 3, 4, 5, 6]
    # k = 7
    # arr = [1, 2, 3, 4, 5, 6]
    # k = 10
    # arr = [-10, 10]
    # k = 2
    arr = [-1, 1, -2, 2, -3, 3, -4, 4]
    k = 3
    print(Solution().canArrange(arr, k))
