from typing import List
from collections import deque


class Solution:
    def maxResult(self, nums: List[int], k: int) -> int:
        queue = deque()
        queue.append(0)
        for i in range(1, len(nums)):
            if queue[0] + k < i:
                queue.popleft()
            nums[i] += nums[queue[0]]
            while len(queue) and nums[queue[-1]] <= nums[i]:
                queue.pop()
            queue.append(i)
        return nums[-1]
