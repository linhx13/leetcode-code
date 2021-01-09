from typing import List
import collections


class Solution:
    def slidingPuzzle(self, board: List[List[int]]) -> int:
        def neighours(state):
            state = list(state)
            idx = state.index("0")
            i0, j0 = idx // 3, idx % 3

            for d in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                i, j = i0 + d[0], j0 + d[1]
                if 0 <= i < 2 and 0 <= j < 3:
                    state[idx], state[i * 3 + j] = state[i * 3 + j], "0"
                    yield "".join(state)
                    state[idx], state[i * 3 + j] = "0", state[idx]

        state = "".join("".join(map(str, row)) for row in board)
        queue = collections.deque([state])
        seen = set([state])
        res = 0
        while len(queue) > 0:
            for _ in range(len(queue)):
                state = queue.popleft()
                if state == "123450":
                    return res
                for s in neighours(state):
                    if s not in seen:
                        queue.append(s)
                        seen.add(s)
            res += 1
        return -1


if __name__ == "__main__":
    # board = [[1, 2, 3], [4, 0, 5]]
    # board = [[1, 2, 3], [5, 4, 0]]
    # board = [[4, 1, 2], [5, 0, 3]]
    board = [[3, 2, 4], [1, 5, 0]]
    print(Solution().slidingPuzzle(board))
