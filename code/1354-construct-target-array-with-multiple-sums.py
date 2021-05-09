from typing import List
import heapq


class Solution:
    def isPossible(self, target: List[int]) -> bool:
        h = []
        total = 0
        for x in target:
            total += x
            heapq.heappush(h, -x)
        while h:
            x = -heapq.heappop(h)
            if x == 1:
                break
            rest = total - x
            if rest <= 0 or x <= rest:
                return False
            total = total - (x // rest * rest)
            x = x % rest
            if x > 1:
                heapq.heappush(h, -x)
        return True


if __name__ == "__main__":
    # target = [9, 3, 5]
    # target = [1, 1, 1, 2]
    # target = [8, 5]
    target = [1, 1000000000]
    # target = [1, 1, 2]
    print(Solution().isPossible(target))
