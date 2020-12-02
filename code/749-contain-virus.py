from typing import List


class Solution:
    def containVirus(self, grid: List[List[int]]) -> int:
        nr, nc = len(grid), len(grid[0])

        def neighbors(r, c):
            for i, j in ((r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1)):
                if 0 <= i < nr and 0 <= j < nc:
                    yield (i, j)

        def dfs(r, c):
            if (r, c) not in seen:
                seen.add((r, c))
                regions[-1].add((r, c))
                for i, j in neighbors(r, c):
                    if grid[i][j] == 1:
                        dfs(i, j)
                    elif grid[i][j] == 0:
                        frontiers[-1].add((i, j))
                        perimeters[-1] += 1

        res = 0
        while True:
            seen = set()
            regions = []
            frontiers = []
            perimeters = []
            for r, row in enumerate(grid):
                for c, v in enumerate(row):
                    if v == 1 and (r, c) not in seen:
                        regions.append(set())
                        frontiers.append(set())
                        perimeters.append(0)
                        dfs(r, c)

            if not regions:
                break

            idx = frontiers.index(max(frontiers, key=len))
            res += perimeters[idx]

            for i, region in enumerate(regions):
                if i == idx:
                    for r, c in region:
                        grid[r][c] = -1
                else:
                    for r, c in region:
                        for i, j in neighbors(r, c):
                            if grid[i][j] == 0:
                                grid[i][j] = 1

        return res
