from typing import List


class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        d = {0: -1}
        if k == 0:
            for i in range(1, len(nums)):
                if nums[i] == 0 and nums[i - 1] == 0:
                    return True
            return False
        s = 0
        for i in range(len(nums)):
            s += nums[i]
            s = s % k
            t = d.get(s)
            if t is not None:
                if i - t >= 2:
                    return True
            else:
                d[s] = i

        return False


if __name__ == "__main__":
    # nums = [23, 2, 4, 6, 7]
    # k = 1
    nums = [23, 2, 6, 4, 7]
    k = 6
    print(Solution().checkSubarraySum(nums, k))
