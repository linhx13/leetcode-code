from typing import List
import collections


class Solution:
    def deleteAndEarn(self, nums: List[int]) -> int:
        if not nums:
            return 0
        counter = collections.Counter(nums)
        yes, no = 0, 0
        prev = None

        for cur in sorted(counter.keys()):
            if cur - 1 != prev:
                yes, no = max(yes, no) + cur * counter[cur], max(yes, no)
            else:
                yes, no = no + cur * counter[cur], max(yes, no)
            prev = cur
        return max(yes, no)


if __name__ == "__main__":
    # nums = [2, 2, 3, 3, 4, 5]
    # nums = [2, 2, 3, 3, 3, 4]
    nums = [3, 4, 2]
    # nums = [1]
    print(Solution().deleteAndEarn(nums))
