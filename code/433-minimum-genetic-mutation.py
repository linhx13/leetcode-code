from typing import List
from collections import deque


class Solution:
    def minMutation(self, start: str, end: str, bank: List[str]) -> int:
        bank = set(bank)
        bank.add(start)
        if end not in bank:
            return -1
        queue = deque()
        visited = set()
        queue.append(start)
        visited.add(start)
        ans = 0
        chars = ['A', 'T', 'C', 'G']
        while len(queue) > 0:
            for _ in range(len(queue)):
                cur = queue.popleft()
                if cur == end:
                    return ans
                for i in range(len(cur)):
                    for c in chars:
                        tmp = cur[:i] + c + cur[i + 1:]
                        if tmp in bank and tmp not in visited:
                            queue.append(tmp)
                            visited.add(tmp)
            ans += 1
        return -1
