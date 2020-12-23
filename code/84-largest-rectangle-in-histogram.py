from typing import List


class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        stack = []
        res = 0
        for i in range(len(heights)):
            if len(stack) == 0 or heights[i] >= heights[stack[-1][0]]:
                stack.append((i, 0))
            else:
                cur_cnt = 0
                while len(stack) > 0 and heights[i] < heights[stack[-1][0]]:
                    old_i, old_cnt = stack.pop()
                    area = heights[old_i] * (old_cnt + i - old_i)
                    res = max(res, area)
                cur_cnt = i if len(stack) == 0 else (i - stack[-1][0] - 1)
                stack.append((i, cur_cnt))
        for i, c in stack:
            res = max(res, heights[i] * (c + stack[-1][0] - i + 1))
        return res


if __name__ == "__main__":
    # heights = [2, 3, 3, 2, 3]
    # heights = [2, 1, 5, 6, 2, 3]
    # heights = [2, 3, 1, 2, 3]
    heights = [3, 6, 5, 7, 4, 8, 1, 0]
    print(Solution().largestRectangleArea(heights))
