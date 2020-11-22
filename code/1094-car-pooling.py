from typing import List


class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        events = []
        for t in trips:
            events.append((t[1], t[0]))
            events.append((t[2], -t[0]))
        events.sort()
        cnt = 0
        for e in events:
            cnt += e[1]
            if cnt > capacity:
                return False
        return True
