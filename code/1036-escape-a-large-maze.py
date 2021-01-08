from typing import List
import collections


class Solution:
    def isEscapePossible(
        self, blocked: List[List[int]], source: List[int], target: List[int]
    ) -> bool:
        if len(blocked) == 0:
            return True

        blocked = set(map(tuple, blocked))

        def bfs(source, target):
            queue = collections.deque([source])
            seen = {tuple(source)}
            level = 0
            while (
                len(queue) > 0
                and level < 2 * len(blocked)
                and len(queue) <= len(blocked)
            ):
                for _ in range(len(queue)):
                    x0, y0 = queue.popleft()
                    for dx, dy in [(0, 1), (0, -1), (-1, 0), (1, 0)]:
                        x, y = x0 + dx, y0 + dy
                        if (
                            0 <= x < 10 ** 6
                            and 0 <= y < 10 ** 6
                            and (x, y) not in seen
                            and (x, y) not in blocked
                        ):
                            if x == target[0] and y == target[1]:
                                return True
                            queue.append((x, y))
                            seen.add((x, y))
                level += 1
            return len(queue) > 0

        return bfs(source, target) and bfs(target, source)


if __name__ == "__main__":
    blocked = [[0, 1], [1, 0]]
    source = [0, 0]
    target = [0, 2]
    # source = [0, 0]
    # target = [999999, 999999]
    print(Solution().isEscapePossible(blocked, source, target))
