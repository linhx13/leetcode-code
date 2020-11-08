from typing import List
from collections import deque


class Solution:
    def canReach(self, arr: List[int], start: int) -> bool:
        queue = deque()
        visited = set()
        queue.append(start)
        visited.add(start)
        while queue:
            idx = queue.popleft()
            if arr[idx] == 0:
                return True
            k = idx + arr[idx]
            if 0 <= k < len(arr) and k not in visited:
                queue.append(k)
                visited.add(k)
            k = idx - arr[idx]
            if 0 <= k < len(arr) and k not in visited:
                queue.append(k)
                visited.add(k)
        return False
