from typing import List
from collections import Counter, defaultdict


class Solution:
    def countSubTrees(
        self, n: int, edges: List[List[int]], labels: str
    ) -> List[int]:
        es = defaultdict(set)
        for a, b in edges:
            es[a].add(b)
            es[b].add(a)
        children = [list() for _ in range(n)]

        def build(i, visited):
            if i in visited:
                return
            visited.add(i)
            for c in es[i]:
                if c not in visited:
                    children[i].append(c)
                    build(c, visited)

        build(0, set())
        res = [0] * n

        def dfs(i):
            counter = Counter({labels[i]: 1})
            for c in children[i]:
                counter.update(dfs(c))
            res[i] = counter[labels[i]]
            return counter

        dfs(0)
        return res


if __name__ == "__main__":
    # n = 7
    # edges = [[0, 1], [0, 2], [1, 4], [1, 5], [2, 3], [2, 6]]
    # labels = "abaedcd"
    # n = 4
    # edges = [[0, 1], [1, 2], [0, 3]]
    # labels = "bbbb"
    # n = 5
    # edges = [[0, 1], [0, 2], [1, 3], [0, 4]]
    # labels = "aabab"
    # n = 6
    # edges = [[0, 1], [0, 2], [1, 3], [3, 4], [4, 5]]
    # labels = "cbabaa"
    # n = 7
    # edges = [[0, 1], [1, 2], [2, 3], [3, 4], [4, 5], [5, 6]]
    # labels = "aaabaaa"
    n = 4
    edges = [[0, 2], [0, 3], [1, 2]]
    labels = "aeed"

    print(Solution().countSubTrees(n, edges, labels))
