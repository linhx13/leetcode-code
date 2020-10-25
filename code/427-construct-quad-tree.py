from typing import List


class Node:
    def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft,
                 bottomRight):
        self.val = val
        self.isLeaf = isLeaf
        self.topLeft = topLeft
        self.topRight = topRight
        self.bottomLeft = bottomLeft
        self.bottomRight = bottomRight


class Solution:
    def construct(self, grid: List[List[int]]) -> 'Node':
        return self.helper(grid, 0, len(grid), 0, len(grid[0]))

    def helper(self, grid, r1, r2, c1, c2):
        if r1 + 1 == r2 and c1 + 1 == c2:
            node = Node(val=bool(grid[r1][c1]),
                        isLeaf=True,
                        topLeft=None,
                        topRight=None,
                        bottomLeft=None,
                        bottomRight=None)
            return node
        top_left = self.helper(grid, r1, r1 + (r2 - r1) // 2, c1,
                               c1 + (c2 - c1) // 2)
        bottom_left = self.helper(grid, r1 + (r2 - r1) // 2, r2, c1,
                                  c1 + (c2 - c1) // 2)
        top_right = self.helper(grid, r1, r1 + (r2 - r1) // 2,
                                c1 + (c2 - c1) // 2, c2)
        bottom_right = self.helper(grid, r1 + (r2 - r1) // 2, r2,
                                   c1 + (c2 - c1) // 2, c2)
        if (top_left.isLeaf and top_right.isLeaf and bottom_left.isLeaf
                and bottom_right.isLeaf and top_left.val == top_right.val
                and top_right.val == bottom_left.val
                and bottom_left.val == bottom_right.val
                and bottom_right.val == top_left.val):
            node = Node(val=top_left.val,
                        isLeaf=True,
                        topLeft=None,
                        topRight=None,
                        bottomLeft=None,
                        bottomRight=None)
        else:
            node = Node(val=True,
                        isLeaf=False,
                        topLeft=top_left,
                        topRight=top_right,
                        bottomLeft=bottom_left,
                        bottomRight=bottom_right)
        return node


if __name__ == "__main__":
    grid = [[1, 1, 1, 1, 0, 0, 0, 0], [1, 1, 1, 1, 0, 0, 0, 0],
            [1, 1, 1, 1, 1, 1, 1, 1], [1, 1, 1, 1, 1, 1, 1, 1],
            [1, 1, 1, 1, 0, 0, 0, 0], [1, 1, 1, 1, 0, 0, 0, 0],
            [1, 1, 1, 1, 0, 0, 0, 0], [1, 1, 1, 1, 0, 0, 0, 0]]
    print(Solution().construct(grid))
