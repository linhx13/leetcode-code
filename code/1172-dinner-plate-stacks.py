import heapq
from typing import List


class DinnerPlates:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.stacks: List[List[int]] = []
        self.pushable_heap: List[int] = []
        self.popable_heap: List[int] = []

    def push(self, val: int) -> None:
        if len(self.pushable_heap) == 0:
            self.stacks.append([])
            idx = len(self.stacks) - 1
        else:
            idx = heapq.heappop(self.pushable_heap)
        self.stacks[idx].append(val)
        if len(self.stacks[idx]) < self.capacity:
            heapq.heappush(self.pushable_heap, idx)
        if len(self.stacks[idx]) == 1:
            heapq.heappush(self.popable_heap, -idx)

    def pop(self) -> int:
        if len(self.popable_heap) == 0:
            return -1
        while True:
            idx = -heapq.heappop(self.popable_heap)
            if len(self.stacks[idx]) > 0:
                break
        val = self.stacks[idx].pop()
        if len(self.stacks[idx]) > 0:
            heapq.heappush(self.popable_heap, -idx)
        if len(self.stacks[idx]) == self.capacity - 1:
            heapq.heappush(self.pushable_heap, idx)
        return val

    def popAtStack(self, index: int) -> int:
        if index >= len(self.stacks) or len(self.stacks[index]) == 0:
            return -1
        val = self.stacks[index].pop()
        if len(self.stacks[index]) == self.capacity - 1:
            heapq.heappush(self.pushable_heap, index)
        return val
