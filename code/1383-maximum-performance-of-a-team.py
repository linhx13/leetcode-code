from typing import List
import heapq


class Solution:
    def maxPerformance(
        self, n: int, speed: List[int], efficiency: List[int], k: int
    ) -> int:
        mod = 10 ** 9 + 7

        candidates = list(zip(efficiency, speed))
        candidates = sorted(candidates, key=lambda t: t[0], reverse=True)

        speed_heap = []
        speed_sum, perf = 0, 0
        for cur_eff, cur_speed in candidates:
            if len(speed_heap) > k - 1:
                speed_sum -= heapq.heappop(speed_heap)
            heapq.heappush(speed_heap, cur_speed)

            speed_sum += cur_speed
            perf = max(perf, speed_sum * cur_eff)
        return perf
