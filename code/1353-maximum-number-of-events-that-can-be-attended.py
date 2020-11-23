from typing import List


class SegmentTreeNode:
    def __init__(self, start, end):
        self.start = start
        self.end = end
        self.left = None
        self.right = None
        self.val = 0


class Solution:
    def maxEvents(self, events: List[List[int]]) -> int:
        if not events:
            return 0
        events.sort(key=lambda x: (x[1], x[0]))
        last_day = events[-1][1]
        first_day = 999999999
        for e in events:
            first_day = min(first_day, e[0])
        root = self.build(first_day, last_day)
        res = 0
        for e in events:
            day = self.query(root, e[0], e[1])
            if day != 999999999:
                res += 1
                self.update(root, day)
        return res

    def build(self, start, end):
        if start > end:
            return None
        node = SegmentTreeNode(start, end)
        node.val = start
        if start != end:
            mid = start + (end - start) // 2
            node.left = self.build(start, mid)
            node.right = self.build(mid + 1, end)
        return node

    def update(self, node, last_day):
        if node.start == node.end:
            node.val = 999999999
        else:
            mid = node.start + (node.end - node.start) // 2
            if mid >= last_day:
                self.update(node.left, last_day)
            else:
                self.update(node.right, last_day)
            node.val = min(node.left.val, node.right.val)

    def query(self, node, start, end):
        if node.start == start and node.end == end:
            return node.val
        mid = node.start + (node.end - node.start) // 2
        if mid >= end:
            return self.query(node.left, start, end)
        elif mid < start:
            return self.query(node.right, start, end)
        else:
            return min(self.query(node.left, start, mid),
                       self.query(node.right, mid + 1, end))


if __name__ == "__main__":
    # events = [[1, 2], [2, 3], [3, 4]]
    # events = [[1, 2], [2, 3], [3, 4], [1, 2]]
    # events = [[1, 4], [4, 4], [2, 2], [3, 4], [1, 1]]
    # events = [[1, 100000]]
    # events = [[1, 1], [1, 2], [1, 3], [1, 4], [1, 5], [1, 6], [1, 7]]
    events = [[1, 2], [1, 2], [3, 3], [1, 5], [1, 5]]
    print(Solution().maxEvents(events))
