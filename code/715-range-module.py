class SegmentTreeNode:
    def __init__(self, start, end, val=False):
        self.start = start
        self.end = end
        self.left = None
        self.right = None
        self.val = val


class RangeModule:
    def __init__(self):
        self.root = SegmentTreeNode(0, 1000000000)

    def addRange(self, left: int, right: int) -> None:
        self._update(self.root, left, right, True)

    def queryRange(self, left: int, right: int) -> bool:
        return self._query(self.root, left, right, True)

    def removeRange(self, left: int, right: int) -> None:
        self._update(self.root, left, right, False)

    def _update(self, node, start, end, val):
        if start <= node.start and node.end <= end:
            node.val = val
            node.left = None
            node.right = None
            return node.val
        if start >= node.end or end <= node.start:
            return node.val
        mid = node.start + (node.end - node.start) // 2
        if node.left is None or node.right is None:
            node.left = SegmentTreeNode(node.start, mid, node.val)
            node.right = SegmentTreeNode(mid, node.end, node.val)
        left_val = self._update(node.left, start, end, val)
        right_val = self._update(node.right, start, end, val)
        node.val = left_val and right_val
        return node.val

    def _query(self, node, start, end, parent_val):
        # if node:
        #     print(start, end, node, node.start, node.end, node.val)
        # else:
        #     print(start, end, node)
        if node is None:
            return parent_val
        if start <= node.start and node.end <= end:
            return node.val
        if start >= node.end or end <= node.start:
            return False
        mid = node.start + (node.end - node.start) // 2
        left_val, right_val = True, True
        if start < mid:
            left_val = self._query(node.left, start, min(mid, end), node.val)
        if end > mid:
            right_val = self._query(node.right, max(start, mid), end, node.val)
        return left_val and right_val


if __name__ == "__main__":
    obj = RangeModule()
    obj.addRange(10, 20)
    obj.removeRange(14, 16)
    print(obj.queryRange(10, 14))
    print(obj.queryRange(13, 15))
    print(obj.queryRange(16, 17))
    # obj.addRange(5, 8)
    # print(obj.queryRange(3, 4))
    # obj.addRange(10, 180)
    # obj.addRange(150, 200)
    # obj.addRange(250, 500)
    # print(obj.queryRange(50, 100))
