from typing import List
from collections import Counter


class Solution:
    def findCenter(self, edges: List[List[int]]) -> int:
        counter = Counter()
        for u, v in edges:
            counter[u] += 1
            counter[v] += 1
            if counter[u] > 1:
                return u
            if counter[v] > 1:
                return v
