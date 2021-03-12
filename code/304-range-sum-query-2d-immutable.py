from typing import List


class NumMatrix:
    def __init__(self, matrix: List[List[int]]):
        self.matrix = matrix
        self.m = len(matrix)
        self.n = 0 if self.m == 0 else len(matrix[0])
        if self.m == 0 or self.n == 0:
            return
        self.m, self.n = len(matrix), len(matrix[0])
        for i in range(self.m):
            for j in range(self.n):
                if i - 1 >= 0:
                    self.matrix[i][j] += self.matrix[i - 1][j]
                if j - 1 >= 0:
                    self.matrix[i][j] += self.matrix[i][j - 1]
                if i - 1 >= 0 and j - 1 >= 0:
                    self.matrix[i][j] -= self.matrix[i - 1][j - 1]

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        if self.m == 0 or self.n == 0:
            return 0
        if row2 < 0 or col2 < 0:
            return 0
        row1, col1 = max(0, row1) - 1, max(0, col1) - 1
        res = self.matrix[row2][col2]
        if row1 >= 0:
            res -= self.matrix[row1][col2]
        if col1 >= 0:
            res -= self.matrix[row2][col1]
        if row1 >= 0 and col1 >= 0:
            res += self.matrix[row1][col1]
        return res
