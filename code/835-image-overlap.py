from typing import List


class Solution:
    def largestOverlap(self, img1: List[List[int]],
                       img2: List[List[int]]) -> int:
        ans = 0
        dim = len(img1)
        for dy in range(0, dim):
            for dx in range(0, dim):
                ans = max(ans, self.shift_and_count(dx, dy, img1, img2))
                ans = max(ans, self.shift_and_count(dx, dy, img2, img1))
        return ans

    def shift_and_count(self, dx, dy, img1, img2):
        dim = len(img1)
        left_cnt, right_cnt = 0, 0
        for row_2, row_1 in enumerate(range(dy, dim)):
            for col_2, col_1 in enumerate(range(dx, dim)):
                if img1[row_1][col_1] == 1 and img1[row_1][col_1] == img2[
                        row_2][col_2]:
                    left_cnt += 1
                if img1[row_1][col_2] == 1 and img1[row_1][col_2] == img2[
                        row_2][col_1]:
                    right_cnt += 1
        return max(left_cnt, right_cnt)
