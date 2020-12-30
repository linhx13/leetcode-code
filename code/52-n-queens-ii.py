class Solution:
    def totalNQueens(self, n: int) -> int:
        def check(row, col):
            return (cols[col] + hill[row - col] + dale[row + col]) == 0

        def place(row, col):
            queens.add((row, col))
            cols[col] = 1
            hill[row - col] = 1
            dale[row + col] = 1

        def remove(row, col):
            queens.remove((row, col))
            cols[col] = 0
            hill[row - col] = 0
            dale[row + col] = 0

        def backtrack(row):
            nonlocal res
            for col in range(n):
                if check(row, col):
                    place(row, col)
                    if row + 1 == n:
                        res += 1
                    else:
                        backtrack(row + 1)
                    remove(row, col)

        cols = [0] * n
        hill = [0] * (2 * n - 1)
        dale = [0] * (2 * n - 1)
        queens = set()
        res = 0
        backtrack(0)
        return res
