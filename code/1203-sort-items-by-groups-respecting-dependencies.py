from typing import List
from collections import defaultdict


class Solution:
    def sortItems(
        self, n: int, m: int, group: List[int], beforeItems: List[List[int]]
    ) -> List[int]:
        group2item = defaultdict(set)
        for i in range(n):
            if group[i] == -1:
                group[i] = m
                m += 1
            group2item[group[i]].add(i)

        item_graph = [set() for _ in range(n)]
        group_graph = [set() for _ in range(m)]
        for i, items in enumerate(beforeItems):
            for j in items:
                item_graph[i].add(j)
                if group[j] != group[i]:
                    group_graph[group[i]].add(group[j])
        group_res = self.topo_sort(group_graph)

        res = []
        for g in group_res:
            if not group2item[g]:
                continue
            items = {}
            rev_items = {}
            for idx, i in enumerate(group2item[g]):
                items[i] = idx
                rev_items[idx] = i
            graph = [set() for _ in range(len(items))]
            for u in group2item[g]:
                for v in item_graph[u]:
                    if v not in group2item[g]:
                        continue
                    graph[items[u]].add(items[v])
            item_res = self.topo_sort(graph)
            if not item_res:
                return []
            for i in item_res:
                res.append(rev_items[i])
        return res

    def topo_sort(self, graph):
        visited = [
            0 for _ in range(len(graph))
        ]  # 0: unvisited, 1: travsed, 2: travlling
        res = []

        def dfs(u):
            if visited[u] == 2:
                return False
            visited[u] = 2
            for v in graph[u]:
                if visited[v] == 0:
                    if not dfs(v):
                        return False
                elif visited[v] == 2:
                    return False
            visited[u] = 1
            res.append(u)
            return True

        for u in range(len(graph)):
            if visited[u] == 0:
                if not dfs(u):
                    return []
        return res


if __name__ == "__main__":
    # n = 8
    # m = 2
    # group = [-1, -1, 1, 0, 0, 1, 0, -1]
    # beforeItems = [[], [6], [5], [6], [3, 6], [], [], []]
    n = 8
    m = 2
    group = [-1, -1, 1, 0, 0, 1, 0, -1]
    beforeItems = [[], [6], [5], [6], [3], [], [4], []]
    print(Solution().sortItems(n, m, group, beforeItems))
