from typing import List


class Solution:
    def sequentialDigits(self, low: int, high: int) -> List[int]:
        res = []

        def helper(n):
            if low <= n <= high:
                res.append(n)
            if n >= high:
                return
            k = n % 10
            if k < 9:
                helper(n * 10 + (k + 1))

        for i in range(1, 10):
            helper(i)
        res.sort()
        return res


if __name__ == "__main__":
    # low = 100
    # high = 300
    # low = 1
    # high = 20
    low = 58
    high = 155
    print(Solution().sequentialDigits(low, high))
