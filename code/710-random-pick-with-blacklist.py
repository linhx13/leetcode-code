from typing import List
import random


class Solution:
    def __init__(self, N: int, blacklist: List[int]):
        blacklist = set(blacklist)
        self.d = {}
        for x in blacklist:
            self.d[x] = -1
        self.m = N - len(blacklist)
        for x in blacklist:
            if x < self.m:
                while (N - 1) in self.d:
                    N -= 1
                self.d[x] = N - 1
                N -= 1

    def pick(self) -> int:
        x = random.randint(0, self.m - 1)
        if x in self.d:
            return self.d[x]
        return x
