from typing import List


class Solution:
    def maxProductPath(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dp_max = [[0 for _ in range(n)] for _ in range(m)]
        dp_min = [[0 for _ in range(n)] for _ in range(m)]
        for i in range(m):
            for j in range(n):
                x = grid[i][j]
                if i == 0 and j == 0:
                    dp_max[i][j] = dp_min[i][j] = x
                    continue
                nums = []
                if j - 1 >= 0:
                    nums.append(x * dp_max[i][j - 1])
                    nums.append(x * dp_min[i][j - 1])
                if i - 1 >= 0:
                    nums.append(x * dp_max[i - 1][j])
                    nums.append(x * dp_min[i - 1][j])
                dp_max[i][j] = max(nums)
                dp_min[i][j] = min(nums)
        if dp_max[m - 1][n - 1] < 0:
            return -1
        else:
            return dp_max[m - 1][n - 1] % (10 ** 9 + 7)


if __name__ == "__main__":
    # grid = [[-1, -2, -3], [-2, -3, -3], [-3, -3, -2]]
    # grid = [[1, -2, 1], [1, -2, 1], [3, -4, 1]]
    # grid = [[1, 3], [0, -4]]
    grid = [[1, 4, 4, 0], [-2, 0, 0, 1], [1, -1, 1, 1]]
    print(Solution().maxProductPath(grid))
