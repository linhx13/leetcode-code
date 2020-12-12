from typing import List


class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        if not nums:
            return []
        res = [-1] * len(nums)
        stack = []
        for _ in range(2):
            for i, x in enumerate(nums):
                while stack and nums[stack[-1]] < x:
                    j = stack.pop()
                    res[j] = x
                stack.append(i)
        return res


if __name__ == "__main__":
    nums = [1, 2, 1, 3]
    print(Solution().nextGreaterElements(nums))
