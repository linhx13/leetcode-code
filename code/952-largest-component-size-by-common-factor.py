from typing import List
import math
from collections import defaultdict, Counter


class UnionFind:
    def __init__(self, n):
        self.p = [i for i in range(n)]

    def find(self, x):
        if self.p[x] != x:
            self.p[x] = self.find(self.p[x])
        return self.p[x]

    def union(self, x, y):
        px, py = self.find(x), self.find(y)
        self.p[px] = py


class Solution:
    def largestComponentSize(self, A: List[int]) -> int:
        n = len(A)
        uf = UnionFind(n)
        primes = defaultdict(list)
        for i, num in enumerate(A):
            prs = self.get_primes(num)
            for p in prs:
                primes[p].append(i)
        for _, nodes in primes.items():
            for i in range(len(nodes) - 1):
                uf.union(nodes[i], nodes[i + 1])
        return max(Counter([uf.find(i) for i in range(n)]).values())

    def get_primes(self, n):
        for i in range(2, int(math.sqrt(n)) + 1):
            if n % i == 0:
                return self.get_primes(n // i) | set([i])
        return set([n])
