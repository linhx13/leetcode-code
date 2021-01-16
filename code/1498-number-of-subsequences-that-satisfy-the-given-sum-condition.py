from typing import List


class Solution:
    def numSubseq(self, nums: List[int], target: int) -> int:
        nums.sort()
        mod = 10 ** 9 + 7
        i, j = 0, len(nums) - 1
        res = 0
        while i <= j:
            if nums[i] + nums[j] > target:
                j -= 1
            else:
                res += pow(2, j - i, mod)
                i += 1
        return res % mod


if __name__ == "__main__":
    # nums = [3, 5, 6, 7]
    # target = 9
    nums = [3, 3, 6, 8]
    target = 10
    print(Solution().numSubseq(nums, target))
