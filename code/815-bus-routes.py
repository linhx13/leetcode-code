from typing import List
from collections import defaultdict, deque


class Solution:
    def numBusesToDestination(self, routes: List[List[int]], S: int,
                              T: int) -> int:
        if S == T:
            return 0
        stop2routes = defaultdict(set)
        for i, r in enumerate(routes):
            for s in r:
                stop2routes[s].add(i)
        graph = defaultdict(set)
        for rs in stop2routes.values():
            for i in rs:
                for j in rs:
                    if i != j:
                        graph[i].add(j)
                        graph[j].add(i)
        starts = stop2routes[S]
        targets = stop2routes[T]
        if not starts or not targets:
            return -1
        queue = deque(starts)
        visited = set(starts)
        res = 1
        while queue:
            for _ in range(len(queue)):
                cur = queue.popleft()
                if cur in targets:
                    return res
                for r in graph[cur]:
                    if r not in visited:
                        queue.append(r)
                        visited.add(r)
            res += 1
        return -1


if __name__ == "__main__":
    routes = [[1, 7], [3, 5]]
    S = 5
    T = 5
    print(Solution().numBusesToDestination(routes, S, T))
