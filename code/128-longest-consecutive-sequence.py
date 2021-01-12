from typing import List


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        nums = set(nums)
        res = 0

        for x in nums:
            if x - 1 not in nums:
                y = x
                c = 1
                while y + 1 in nums:
                    y += 1
                    c += 1
                res = max(res, c)
        return res


if __name__ == "__main__":
    # nums = [100, 4, 200, 1, 3, 2]
    nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
    print(Solution().longestConsecutive(nums))
