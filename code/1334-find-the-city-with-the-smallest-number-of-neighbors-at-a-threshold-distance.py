from typing import List


class Solution:
    def findTheCity(self, n: int, edges: List[List[int]],
                    distanceThreshold: int) -> int:
        d = [[999999999 for _ in range(n)] for _ in range(n)]
        for i in range(n):
            d[i][i] = 0
        for e in edges:
            d[e[0]][e[1]] = min(d[e[0]][e[1]], e[2])
            d[e[1]][e[0]] = min(d[e[1]][e[0]], e[2])

        for k in range(n):
            for i in range(n):
                for j in range(n):
                    if d[i][j] > d[i][k] + d[k][j]:
                        d[i][j] = d[i][k] + d[k][j]

        ans = -1
        min_cnt = 9999999999
        for i in range(n):
            cnt = 0
            for j in range(n):
                if d[i][j] <= distanceThreshold:
                    cnt += 1
            if cnt <= min_cnt:
                min_cnt = cnt
                ans = i
        return ans
