from typing import List


class Solution:
    def numSubarrayBoundedMax(self, A: List[int], L: int, R: int) -> int:
        res = [0]
        pre = -1
        for i in range(len(A)):
            if A[i] < L:
                res.append(res[-1])
            elif A[i] > R:
                res.append(0)
                pre = i
            else:
                res.append(i - pre)
        return sum(res)


if __name__ == "__main__":
    A = [2, 1, 4, 3]
    L = 2
    R = 3
    print(Solution().numSubarrayBoundedMax(A, L, R))
