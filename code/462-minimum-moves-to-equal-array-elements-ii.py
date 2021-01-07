from typing import List


class Solution:
    def minMoves2(self, nums: List[int]) -> int:
        median = sorted(nums)[len(nums) // 2]
        return sum(abs(x - median) for x in nums)


if __name__ == "__main__":
    nums = [1, 0, 0, 8, 6]
    print(Solution().minMoves2(nums))
