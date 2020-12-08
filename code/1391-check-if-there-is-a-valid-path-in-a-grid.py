from typing import List


class Solution:
    def hasValidPath(self, grid: List[List[int]]) -> bool:
        if not grid:
            return True
        directions = {
            1: [(0, -1), (0, 1)],
            2: [(-1, 0), (1, 0)],
            3: [(0, -1), (1, 0)],
            4: [(0, 1), (1, 0)],
            5: [(0, -1), (-1, 0)],
            6: [(0, 1), (-1, 0)],
        }
        visited = set()
        m, n = len(grid), len(grid[0])

        def dfs(i, j):
            visited.add((i, j))
            if not (0 <= i < m and 0 <= j < n):
                return False
            if i == m - 1 and j == n - 1:
                return True
            for d in directions[grid[i][j]]:
                ii, jj = i + d[0], j + d[1]
                if (
                    0 <= ii < m
                    and 0 <= jj < n
                    and (ii, jj) not in visited
                    and (-d[0], -d[1]) in directions[grid[ii][jj]]
                ):
                    if dfs(ii, jj):
                        return True
            return False

        return dfs(0, 0)


if __name__ == "__main__":
    grid = [[1, 1, 2]]
    print(Solution().hasValidPath(grid))
