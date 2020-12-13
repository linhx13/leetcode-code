from typing import List
from collections import defaultdict, deque


class Solution:
    def minJumps(self, arr: List[int]) -> int:
        n = len(arr)
        g = defaultdict(list)
        for i, x in enumerate(arr):
            g[x].append(i)

        queue = deque()
        visited = set()
        queue.append(0)
        visited.add(0)
        res = 0
        while queue:
            for _ in range(len(queue)):
                cur = queue.popleft()
                if cur == n - 1:
                    return res
                for x in g[arr[cur]]:
                    if x not in visited:
                        queue.append(x)
                        visited.add(x)
                g[arr[cur]].clear()
                for x in [cur - 1, cur + 1]:
                    if 0 <= x < n and x not in visited:
                        queue.append(x)
                        visited.add(x)
            res += 1
        return -1


if __name__ == "__main__":
    arr = [100, -23, -23, 404, 100, 23, 23, 23, 3, 404]
    print(Solution().minJumps(arr))
