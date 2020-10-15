from typing import List


class Solution:
    def findDiagonalOrder(self, nums: List[List[int]]) -> List[int]:
        arr = []
        for i, row in enumerate(nums):
            for j, v in enumerate(row):
                arr.append((i + j, -i, v))
        arr.sort()
        res = [x[2] for x in arr]
        return res


if __name__ == "__main__":
    # nums = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    # nums = [[1, 2, 3, 4, 5], [6, 7], [8], [9, 10, 11], [12, 13, 14, 15, 16]]
    # nums = [[1, 2, 3], [4], [5, 6, 7], [8], [9, 10, 11]]
    nums = [[1, 2, 3, 4, 5, 6]]
    sol = Solution()
    print(sol.findDiagonalOrder(nums))
