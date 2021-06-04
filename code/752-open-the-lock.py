from typing import List
from collections import deque


class Solution:
    def openLock(self, deadends: List[str], target: str) -> int:
        deadends = set(deadends)
        visited = set()
        if "0000" in deadends:
            return -1
        if target == "0000":
            return 0
        queue = deque()
        queue.append("0000")
        visited = set()
        visited.add("0000")
        res = 0
        while len(queue) > 0:
            for _ in range(len(queue)):
                cur = queue.popleft()
                if cur == target:
                    return res
                for i in range(4):
                    for k in [1, -1]:
                        s = list(cur)
                        s[i] = chr(
                            (ord(s[i]) - ord("0") + k + 10) % 10 + ord("0")
                        )
                        s = "".join(s)
                        if s not in visited and s not in deadends:
                            queue.append(s)
                            visited.add(s)
            res += 1
        return -1


if __name__ == "__main__":
    deadends = ["0201", "0101", "0102", "1212", "2002"]
    target = "0202"
    print(Solution().openLock(deadends, target))
