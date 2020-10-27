from typing import List


class Solution:
    def longestOnes(self, A: List[int], K: int) -> int:
        res = 0
        left = 0
        for right in range(len(A)):
            if A[right] == 0:
                K -= 1
            while K < 0:
                if A[left] == 0:
                    K += 1
                left += 1
            res = max(res, right - left + 1)
        return res


if __name__ == "__main__":
    A = [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0]
    # A = [0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1]
    # A = [1, 1, 1, 1, 1]
    # A = [0, 0, 0, 1]
    # A = [0, 0, 0, 0]
    # K = 0
    K = 2
    print(Solution().longestOnes(A, K))
