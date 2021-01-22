from typing import List


class Solution:
    def maxRotateFunction(self, A: List[int]) -> int:
        s = sum(A)
        N = len(A)
        res = f = sum(i * x for i, x in enumerate(A))
        for i in range(1, N):
            f = f + s - N * A[N - i]
            res = max(res, f)
        return res


if __name__ == "__main__":
    A = [4, 3, 2, 6]
    print(Solution().maxRotateFunction(A))
