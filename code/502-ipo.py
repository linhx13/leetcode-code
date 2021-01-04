from typing import List
import heapq


class Solution:
    def findMaximizedCapital(
        self, k: int, W: int, Profits: List[int], Capital: List[int]
    ) -> int:
        avail = []
        other = []
        for p, c in zip(Profits, Capital):
            if c <= W:
                heapq.heappush(avail, -p)
            else:
                heapq.heappush(other, (c, p))
        other.sort(key=lambda x: x[1])
        while k > 0 and len(avail) > 0:
            W += -heapq.heappop(avail)
            k -= 1
            while len(other) > 0 and other[0][0] <= W:
                x = heapq.heappop(other)
                heapq.heappush(avail, -x[1])
        return W


if __name__ == "__main__":
    # k = 2
    # W = 0
    # Profits = [1, 2, 3]
    # Capital = [0, 1, 1]
    k = 11
    W = 11
    Profits = [1, 2, 3]
    Capital = [11, 12, 13]
    print(Solution().findMaximizedCapital(k, W, Profits, Capital))
