from typing import List
import heapq


class Solution:
    def nthSuperUglyNumber(self, n: int, primes: List[int]) -> int:
        s = set()
        h = []
        heapq.heappush(h, 1)
        while n > 0:
            c = heapq.heappop(h)
            for p in primes:
                x = c * p
                if x not in s:
                    heapq.heappush(h, x)
                    s.add(x)
            n -= 1
        return c


if __name__ == "__main__":
    n = 12
    primes = [2, 7, 13, 19]
    sol = Solution()
    print(sol.nthSuperUglyNumber(n, primes))
