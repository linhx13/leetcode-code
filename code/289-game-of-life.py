from typing import List


class Solution:
    def gameOfLife(self, board: List[List[int]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        neighbors = [(1, 0), (1, -1), (0, -1), (-1, -1), (-1, 0), (-1, 1),
                     (0, 1), (1, 1)]
        rows = len(board)
        cols = len(board[0])

        tmp_board = [[board[r][c] for c in range(cols)] for r in range(rows)]

        for row in range(rows):
            for col in range(cols):
                lives = 0
                for n in neighbors:
                    r = row + n[0]
                    c = col + n[1]

                    if 0 <= r < rows and 0 <= c < cols and tmp_board[r][c] == 1:
                        lives += 1
                if tmp_board[row][col] == 1 and (lives < 2 or lives > 3):
                    board[row][col] = 0
                if tmp_board[row][col] == 0 and lives == 3:
                    board[row][col] = 1
