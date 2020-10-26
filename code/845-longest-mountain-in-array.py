from typing import List


class Solution:
    def longestMountain(self, A: List[int]) -> int:
        res = 0
        i = 0
        while i < len(A):
            j = i
            while j + 1 < len(A) and A[j + 1] > A[j]:
                j += 1
            k = j
            while j + 1 < len(A) and A[j + 1] < A[j]:
                j += 1
            if i < k and k < j:
                res = max(j - i + 1, res)
                i = j
            else:
                if i == j:
                    i += 1
                else:
                    i = j
        return res


if __name__ == "__main__":
    # A = [2, 1, 4, 7, 3, 2, 5]
    A = [2, 2, 2]
    print(Solution().longestMountain(A))
