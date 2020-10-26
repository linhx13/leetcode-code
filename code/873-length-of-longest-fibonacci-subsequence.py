from typing import List


class Solution:
    def lenLongestFibSubseq(self, A: List[int]) -> int:
        dp = [[0 for _ in range(len(A))] for _ in range(len(A))]
        m = {}
        for i in range(len(A)):
            m[A[i]] = i
        res = 0
        for i in range(2, len(A)):
            for j in range(i - 1, 0, -1):
                x = A[i] - A[j]
                idx = m.get(x, -1)
                if idx != -1 and idx < j:
                    dp[i][j] = dp[j][idx] + 1
                    res = max(res, dp[i][j])
        if res > 0:
            return res + 2
        else:
            return res


if __name__ == '__main__':
    # A = [1, 2, 3, 4, 5, 6, 7, 8]
    # A = [1, 3, 7, 11, 12, 14, 18]
    A = [1, 2]
    print(Solution().lenLongestFibSubseq(A))
