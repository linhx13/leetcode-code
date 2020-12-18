from typing import List


class Solution:
    def videoStitching(self, clips: List[List[int]], T: int) -> int:
        clips.sort()
        dp = [101 for _ in range(101)]
        dp[0] = 0
        for c in clips:
            for i in range(c[0] + 1, c[1] + 1):
                dp[i] = min(dp[i], dp[c[0]] + 1)
        return -1 if dp[T] >= 100 else dp[T]
