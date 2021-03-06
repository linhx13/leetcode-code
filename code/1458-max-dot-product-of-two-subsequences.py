from typing import List


class Solution:
    def maxDotProduct(self, nums1: List[int], nums2: List[int]) -> int:
        l1, l2 = len(nums1), len(nums2)
        dp = [[float('-inf') for _ in range(l2 + 1)] for _ in range(l1 + 1)]
        for i in range(1, l1 + 1):
            for j in range(1, l2 + 1):
                dp[i][j] = max(
                    nums1[i - 1] * nums2[j - 1],
                    nums1[i - 1] * nums2[j - 1] + dp[i - 1][j - 1],
                    dp[i - 1][j],
                    dp[i][j - 1],
                )
        return dp[l1][l2]


if __name__ == "__main__":
    # nums1 = [2, 1, -2, 5]
    # nums2 = [3, 0, -6]
    # nums1 = [3, -2]
    # nums2 = [2, -6, 7]
    nums1 = [-1, -1]
    nums2 = [1, 1]
    print(Solution().maxDotProduct(nums1, nums2))
