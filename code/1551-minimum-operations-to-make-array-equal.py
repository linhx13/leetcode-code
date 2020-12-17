class Solution:
    def minOperations(self, n: int) -> int:
        last = 2 * (n - 1) + 1
        target = (1 + last) // 2
        c = n // 2
        mid = 2 * (c - 1) + 1
        return c * (target - mid + target - 1) // 2
