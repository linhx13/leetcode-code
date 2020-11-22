from typing import List
from collections import deque


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        queue = deque()
        for i in range(k - 1):
            while queue and queue[-1][0] <= nums[i]:
                queue.pop()
            queue.append((nums[i], i))
        res = []
        for i in range(k - 1, len(nums)):
            while queue and queue[0][1] <= i - k:
                queue.popleft()
            while queue and queue[-1][0] <= nums[i]:
                queue.pop()
            queue.append((nums[i], i))
            res.append(queue[0][0])
        return res
