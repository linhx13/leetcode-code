from typing import List
import bisect
import math


class Node:
    def __init__(self, val, freq):
        self.val = val
        self.freq = freq


def merge(a, b):
    if a.val == b.val:
        return Node(a.val, a.freq + b.freq)
    if a.freq > b.freq:
        return Node(a.val, a.freq - b.freq)
    else:
        return Node(b.val, b.freq - a.freq)


class MajorityChecker:
    def __init__(self, arr: List[int]):
        self.arr = arr
        self.n = len(arr)
        self.size = 1 << (math.ceil(math.log(self.n, 2)) + 1)
        self.tree = [None] * self.size
        self.index = {i: [] for i in self.arr}
        self.build(0, 0, self.n - 1)

    def build(self, pos, l, r):
        if l == r:
            self.tree[pos] = Node(self.arr[l], 1)
            self.index[self.arr[l]].append(l)
        else:
            mid = (l + r) >> 1
            self.build(pos * 2 + 1, l, mid)
            self.build(pos * 2 + 2, mid + 1, r)
            self.tree[pos] = merge(
                self.tree[pos * 2 + 1], self.tree[pos * 2 + 2]
            )

    def cquery(self, pos, start, end, l, r):
        if l > end or r < start:
            return Node(0, 0)
        if start <= l and r <= end:
            return self.tree[pos]
        mid = (l + r) >> 1
        a = self.cquery(pos * 2 + 1, start, end, l, mid)
        b = self.cquery(pos * 2 + 2, start, end, mid + 1, r)
        return merge(a, b)

    def query(self, left: int, right: int, threshold: int) -> int:
        c = self.cquery(0, left, right, 0, self.n - 1).val
        if c == 0:
            return -1
        s = bisect.bisect_left(self.index[c], left)
        e = bisect.bisect_right(self.index[c], right)
        return c if e - s >= threshold else -1
