class SegmentTreeNode:
    def __init__(self, start, end):
        self.start = start
        self.end = end
        self.val = 0
        self.lazy = 0
        self.left = None
        self.right = None

    def get_mid(self):
        return self.start + (self.end - self.start) // 2

    def get_left(self):
        if self.left is None:
            self.left = SegmentTreeNode(self.start, self.get_mid())
        return self.left

    def get_right(self):
        if self.right is None:
            self.right = SegmentTreeNode(self.get_mid() + 1, self.end)
        return self.right


class MyCalendarThree:
    def __init__(self):
        self.root = SegmentTreeNode(0, 1000000000)

    def book(self, start: int, end: int) -> int:
        self._update(self.root, start, end - 1)
        return self.root.val

    def _update(self, root: SegmentTreeNode, start: int, end: int):
        if root.lazy > 0:
            root.val += root.lazy
            if root.start < root.end:
                root.get_left().lazy += root.lazy
                root.get_right().lazy += root.lazy
            root.lazy = 0
        if start <= root.start and root.end <= end:
            root.val += 1
            if root.start < root.end:
                root.get_left().lazy += 1
                root.get_right().lazy += 1
            return
        mid = root.get_mid()
        if mid >= end:
            self._update(root.get_left(), start, end)
        elif mid < start:
            self._update(root.get_right(), start, end)
        else:
            self._update(root.get_left(), start, mid)
            self._update(root.get_right(), mid + 1, end)
        root.val = max(root.val, root.get_left().val, root.get_right().val)


if __name__ == "__main__":
    obj = MyCalendarThree()
    print(obj.book(10, 20))
    print(obj.book(50, 60))
    print(obj.book(10, 40))
    print(obj.book(5, 15))
    print(obj.book(5, 10))
    print(obj.book(25, 55))
