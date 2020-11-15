import heapq
from collections import Counter


class Solution:
    def reorganizeString(self, S: str) -> str:
        counter = Counter(S)
        pq = [(-x[1], x[0]) for x in counter.items()]
        heapq.heapify(pq)
        if any(-c > (len(S) + 1) // 2 for c, _ in pq):
            return ""

        res = []
        while len(pq) >= 2:
            n1, c1 = heapq.heappop(pq)
            n2, c2 = heapq.heappop(pq)
            if not res or c1 != res[-1]:
                res.extend([c1, c2])
            else:
                res.extend([c2, c1])
            if n1 + 1 != 0:
                heapq.heappush(pq, (n1 + 1, c1))
            if n2 + 1 != 0:
                heapq.heappush(pq, (n2 + 1, c2))
        if pq:
            res.append(pq[0][1])
        return "".join(res)
