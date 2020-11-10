from typing import List
from collections import defaultdict


class Solution:
    def canReorderDoubled(self, A: List[int]) -> bool:
        A = sorted(A)
        dd = defaultdict(list)
        for i, x in enumerate(A):
            dd[x].append(i)
        used = [False for _ in range(len(A))]
        cnt = 0
        for i, x in enumerate(A):
            if used[i]:
                continue
            if x >= 0:
                y = 2 * x
            else:
                if x % 2 != 0:
                    cnt += 1
                    continue
                y = x / 2
            if not dd.get(y):
                cnt += 1
                continue
            j = dd[y].pop()
            used[j] = True
        return cnt <= 1


if __name__ == "__main__":
    A = [-2, -6, -3, 4, -4, 2]
    print(Solution().canReorderDoubled(A))
