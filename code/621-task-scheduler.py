from typing import List
from collections import Counter


class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        counter = Counter(tasks)
        max_freq = max(counter.values())

        min_time = -1 + max_freq + (max_freq - 1) * n
        cur_time = 0
        for k in counter:
            if counter[k] == max_freq:
                min_time += 1
            cur_time += counter[k]
        return max(min_time, cur_time)
