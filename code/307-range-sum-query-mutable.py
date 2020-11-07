from typing import List


class SegmentTreeNode:
    def __init__(self, start, end):
        self.start = start
        self.end = end
        self.val = 0
        self.left = None
        self.right = None


class NumArray:
    def __init__(self, nums: List[int]):
        self.root = None
        if len(nums) > 0:
            self.root = self._build_tree(0, len(nums) - 1, nums)

    def update(self, i: int, val: int) -> None:
        if self.root is None:
            return
        self._update_tree(self.root, i, val)

    def sumRange(self, i: int, j: int) -> int:
        if self.root is None:
            return 0
        return self._query_tree(self.root, i, j)

    def _build_tree(self, start, end, nums):
        if start > end:
            return None
        root = SegmentTreeNode(start, end)
        if start == end:
            root.val = nums[start]
            return root
        mid = start + (end - start) // 2
        root.left = self._build_tree(start, mid, nums)
        root.right = self._build_tree(mid + 1, end, nums)
        root.val = root.left.val + root.right.val
        return root

    def _query_tree(self, root, start, end):
        if root.start == start and root.end == end:
            return root.val
        mid = root.start + (root.end - root.start) // 2
        left_val, right_val = 0, 0
        if start <= mid:
            left_val = self._query_tree(root.left, start, min(mid, end))
        if end > mid:
            right_val = self._query_tree(root.right, max(mid + 1, start), end)
        return left_val + right_val

    def _update_tree(self, root, idx, val):
        if root.start == idx and root.end == idx:
            root.val = val
            return
        mid = root.start + (root.end - root.start) // 2
        if idx <= mid:
            self._update_tree(root.left, idx, val)
        else:
            self._update_tree(root.right, idx, val)
        root.val = root.left.val + root.right.val


if __name__ == "__main__":
    nums = [0, 9, 5, 7, 3]
    obj = NumArray(nums)
    print(obj.sumRange(4, 4))
