from typing import List


class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        rows = [set() for _ in range(9)]
        cols = [set() for _ in range(9)]
        boxes = [set() for _ in range(9)]

        def box(r, c):
            return (r // 3) * 3 + (c // 3)

        idx_list = []
        for r in range(9):
            for c in range(9):
                if board[r][c] == '.':
                    idx_list.append((r, c))
                else:
                    x = int(board[r][c])
                    rows[r].add(x)
                    cols[c].add(x)
                    boxes[box(r, c)].add(x)

        def dfs(idx):
            if idx == len(idx_list):
                return True
            r, c = idx_list[idx]
            for x in range(1, 10):
                b = box(r, c)
                if x not in rows[r] and x not in cols[c] and x not in boxes[b]:
                    rows[r].add(x)
                    cols[c].add(x)
                    boxes[b].add(x)
                    board[r][c] = str(x)
                    if dfs(idx + 1):
                        return True
                    else:
                        rows[r].remove(x)
                        cols[c].remove(x)
                        boxes[b].remove(x)
                        board[r][c] = '.'
            return False

        dfs(0)
