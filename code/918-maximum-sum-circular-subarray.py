from typing import List


class Solution:
    def maxSubarraySumCircular(self, A: List[int]) -> int:
        total, max_sum, cur_max, min_sum, cur_min = 0, A[0], 0, A[0], 0
        for x in A:
            cur_max = max(cur_max + x, x)
            max_sum = max(max_sum, cur_max)
            cur_min = min(cur_min + x, x)
            min_sum = min(min_sum, cur_min)
            total += x
        return max(max_sum, total - min_sum) if max_sum > 0 else max_sum


if __name__ == "__main__":
    A = [1, -2, 3, -2]
    # A = [5, -3, 5]
    print(Solution().maxSubarraySumCircular(A))
