from typing import List
import functools


class Solution:
    def PredictTheWinner(self, nums: List[int]) -> bool:
        @functools.lru_cache(None)
        def helper(nums, p1, p2, is_p1):
            if len(nums) == 0:
                if is_p1:
                    return p1 >= p2
                else:
                    return not (p1 >= p2)
            idx_list = [0]
            if len(nums) > 1:
                idx_list.append(len(nums) - 1)
            if is_p1:
                return any(not helper(nums[:i] + nums[i + 1:], p1 +
                                      nums[i], p2, not is_p1)
                           for i in idx_list)
            else:
                return any(not helper(nums[:i] + nums[i + 1:], p1, p2 +
                                      nums[i], not is_p1) for i in idx_list)

        return helper(tuple(nums), 0, 0, True)


if __name__ == '__main__':
    # nums = [1, 5, 2]
    nums = [1, 5, 233, 7]
    print(Solution().PredictTheWinner(nums))
