from typing import List


class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        res = [[0 for _ in range(n)] for _ in range(n)]
        r, c = 0, 0
        x = 1
        while n > 0:
            i, j = r, c
            for j in range(c, c + n):
                res[r][j] = x
                x += 1
            for i in range(r + 1, r + n):
                res[i][c + n - 1] = x
                x += 1
            for j in range(c + n - 2, c - 1, -1):
                res[r + n - 1][j] = x
                x += 1
            for i in range(r + n - 2, r, -1):
                res[i][c] = x
                x += 1
            n -= 2
            r, c = r + 1, c + 1
        return res


if __name__ == "__main__":
    n = 1
    print(Solution().generateMatrix(n))
