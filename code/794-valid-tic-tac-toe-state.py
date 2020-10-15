from typing import List


class Solution:
    def validTicTacToe(self, board: List[str]) -> bool:
        cnt1, cnt2 = 0, 0
        rows = [0, 0, 0]
        cols = [0, 0, 0]
        xrow, xcol, xdia = 0, 0, 0
        orow, ocol, odia = 0, 0, 0
        for x, row in enumerate(board):
            for y, c in enumerate(row):
                if c == "X":
                    cnt1 += 1
                    rows[x] += 1
                    cols[y] += 1
                elif c == "O":
                    cnt2 += 1
                    rows[x] -= 1
                    cols[y] -= 1
                if rows[x] == 3:
                    xrow += 1
                if cols[y] == 3:
                    xcol += 1
                if rows[x] == -3:
                    orow += 1
                if cols[y] == -3:
                    ocol += 1
        if not (0 <= cnt1 - cnt2 <= 1):
            return False
        if board[0][0] == board[1][1] and board[1][1] == board[2][2]:
            if board[0][0] == "X":
                xdia += 1
            elif board[0][0] == "O":
                odia += 1
        if board[0][2] == board[1][1] and board[1][1] == board[2][0]:
            if board[0][2] == "X":
                xdia += 1
            elif board[0][2] == "O":
                odia += 1
        if xrow > 1 or xcol > 1:
            return False
        if orow > 1 or ocol > 1:
            return False
        if (xrow > 0 and orow > 0) or (xcol > 0 and ocol > 0):
            return False
        if (xrow > 0 or xcol > 0 or xdia > 0) and cnt2 == cnt1:
            return False
        if (orow > 0 or ocol > 0 or odia > 0) and cnt1 > cnt2:
            return False
        return True


if __name__ == "__main__":
    # board = ["XXX", "   ", "OOO"]
    # board = ["XOX", " X ", "   "]
    # board = ["XOX", "OXO", "XOX"]
    # board = ["XXX", "XOO", "OO "]
    # board = ["OXX", "XOX", "OXO"]
    board = ["X  ", "   ", "  O"]
    sol = Solution()
    print(sol.validTicTacToe(board))
