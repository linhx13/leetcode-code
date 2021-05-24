from typing import List


class Solution:
    def orderOfLargestPlusSign(self, n: int, mines: List[List[int]]) -> int:
        banned = {tuple(mine) for mine in mines}
        dp = [[0] * n for _ in range(n)]

        for r in range(n):
            cnt = 0
            for c in range(n):
                cnt = 0 if (r, c) in banned else cnt + 1
                dp[r][c] = cnt

            cnt = 0
            for c in range(n - 1, -1, -1):
                cnt = 0 if (r, c) in banned else cnt + 1
                if cnt < dp[r][c]:
                    dp[r][c] = cnt

        res = 0
        for c in range(n):
            cnt = 0
            for r in range(n):
                cnt = 0 if (r, c) in banned else cnt + 1
                if cnt < dp[r][c]:
                    dp[r][c] = cnt

            cnt = 0
            for r in range(n - 1, -1, -1):
                cnt = 0 if (r, c) in banned else cnt + 1
                if cnt < dp[r][c]:
                    dp[r][c] = cnt
                res = max(res, dp[r][c])

        return res
