from typing import List


class Solution:
    def singleNumber(self, nums: List[int]) -> List[int]:
        val = 0
        for n in nums:
            val ^= n
        diff = val & (-val)

        ans = [0, 0]
        for n in nums:
            if n & diff == 0:
                ans[0] ^= n
            else:
                ans[1] ^= n
        return ans
