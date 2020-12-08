from collections import deque


class Solution:
    def findLexSmallestString(self, s: str, a: int, b: int) -> str:
        visited = set()
        queue = deque()
        res = s
        visited.add(s)
        queue.append(s)
        while queue:
            x = queue.popleft()
            res = min(res, x)

            l = list(x)
            for i in range(1, len(l), 2):
                l[i] = str((int(l[i]) + a) % 10)
            t = "".join(l)
            if t not in visited:
                visited.add(t)
                queue.append(t)

            l = list(x)
            for i in range(len(l)):
                j = (i + b) % len(l)
                l[j] = x[i]
            t = "".join(l)
            if t not in visited:
                visited.add(t)
                queue.append(t)
        return res
