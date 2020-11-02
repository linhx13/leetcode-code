from typing import List
from collections import Counter


class Window:
    def __init__(self):
        self.counter = Counter()
        self.nonzero = 0

    def add(self, n):
        self.counter[n] += 1
        if self.counter[n] == 1:
            self.nonzero += 1

    def remove(self, n):
        self.counter[n] -= 1
        if self.counter[n] == 0:
            self.nonzero -= 1


class Solution:
    def subarraysWithKDistinct(self, A: List[int], K: int) -> int:
        window1 = Window()
        window2 = Window()
        ans = 0
        left1 = left2 = 0
        for n in A:
            window1.add(n)
            window2.add(n)

            while window1.nonzero > K:
                window1.remove(A[left1])
                left1 += 1
            while window2.nonzero >= K:
                window2.remove(A[left2])
                left2 += 1
            ans += left2 - left1

        return ans
