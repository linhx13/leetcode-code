from typing import List


class Solution:
    def kConcatenationMaxSum(self, arr: List[int], k: int) -> int:
        n = len(arr)
        s = arr[0]
        m = arr[0]
        mod = 10 ** 9 + 7
        for i in range(1, n * min(k, 2)):
            s = max(arr[i % n], s + arr[i % n])
            m = max(m, s)
        return max(0, m, sum(arr) * max(0, k - 2) + m) % mod
