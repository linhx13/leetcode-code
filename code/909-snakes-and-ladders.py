from typing import List
import collections


class Solution:
    def snakesAndLadders(self, board: List[List[int]]) -> int:
        N = len(board)
        queue = collections.deque()
        seen = set()
        res = 0
        queue.append(1)
        seen.add(1)

        def get_ij(x):
            x -= 1
            i = N - 1 - x // N
            ii = N - 1 - i
            if (N - 1 - i) % 2 == 0:
                j = x % N
            else:
                j = N - 1 - x % N
            return i, j

        while len(queue) > 0:
            for _ in range(len(queue)):
                x = queue.popleft()
                if x == N * N:
                    return res
                for k in range(1, 6 + 1):
                    y = x + k
                    if y > N * N:
                        break
                    i, j = get_ij(y)
                    if board[i][j] != -1:
                        y = board[i][j]
                    if y not in seen:
                        queue.append(y)
                        seen.add(y)
            res += 1
        return -1


if __name__ == "__main__":
    board = [
        [-1, -1, -1, -1, -1, -1],
        [-1, -1, -1, -1, -1, -1],
        [-1, -1, -1, -1, -1, -1],
        [-1, 35, -1, -1, 13, -1],
        [-1, -1, -1, -1, -1, -1],
        [-1, 15, -1, -1, -1, -1],
    ]
    # board = [[-1, -1], [-1, 3]]
    # board = [[-1, 4, -1], [6, 2, 6], [-1, 3, -1]]
    board = [[-1, 1, 2, -1], [2, 13, 15, -1], [-1, 10, -1, -1], [-1, 6, 2, 8]]

    print(Solution().snakesAndLadders(board))
