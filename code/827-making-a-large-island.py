from typing import List


class Solution:
    def largestIsland(self, grid: List[List[int]]) -> int:
        N = len(grid)

        def neighbors(r, c):
            for nr, nc in ((r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1)):
                if 0 <= nr < N and 0 <= nc < N:
                    yield nr, nc

        def dfs(r, c, idx):
            res = 1
            grid[r][c] = idx
            for nr, nc in neighbors(r, c):
                if grid[nr][nc] == 1:
                    res += dfs(nr, nc, idx)
            return res

        area = {}
        idx = 2
        for r in range(N):
            for c in range(N):
                if grid[r][c] == 1:
                    area[idx] = dfs(r, c, idx)
                    idx += 1

        ans = max(area.values() or [0])
        for r in range(N):
            for c in range(N):
                if grid[r][c] == 0:
                    seen = {
                        grid[nr][nc]
                        for nr, nc in neighbors(r, c) if grid[nr][nc] > 1
                    }
                    ans = max(ans, 1 + sum(area[x] for x in seen))
        return ans
