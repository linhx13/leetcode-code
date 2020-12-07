from typing import List


class Solution:
    def largest1BorderedSquare(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        ones = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                ones[i][j] = (
                    ones[i - 1][j]
                    + ones[i][j - 1]
                    - ones[i - 1][j - 1]
                    + grid[i - 1][j - 1]
                )

        def get_ones(i1, j1, i2, j2):
            return ones[i2][j2] - ones[i1][j2] - ones[i2][j1] + ones[i1][j1]

        def check(i, j, l):
            if i < l or j < l:
                return False
            left = get_ones(i - l, j - l, i, j - l + 1)
            right = get_ones(i - l, j - 1, i, j)
            up = get_ones(i - l, j - l, i - l + 1, j)
            down = get_ones(i - 1, j - l, i, j)
            if not (left == l and right == l and up == l and down == l):
                return False
            else:
                return True

        res = 0
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                for l in range(1, min(m, n) + 1):
                    if check(i, j, l):
                        res = max(res, l * l)
        return res


if __name__ == "__main__":
    # grid = [[1, 1, 1], [1, 0, 1], [1, 1, 1]]
    # grid = [[1, 1, 0, 0]]
    grid = [[1, 1, 0], [1, 1, 1], [1, 1, 1], [1, 1, 1]]
    print(Solution().largest1BorderedSquare(grid))
