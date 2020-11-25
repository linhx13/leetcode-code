from typing import List


class Solution:
    def numSimilarGroups(self, strs: List[str]) -> int:
        n = len(strs)
        if n < 2:
            return n

        visited = [False for _ in range(n)]

        def dfs(i):
            visited[i] = True
            for j in range(n):
                if visited[j]:
                    continue
                if similar(strs[i], strs[j]):
                    dfs(j)

        def similar(w1, w2):
            diff = 0
            for x, y in zip(w1, w2):
                if x != y:
                    diff += 1
            return diff <= 2

        res = 0
        for i in range(n):
            if not visited[i]:
                res += 1
                dfs(i)
        return res
