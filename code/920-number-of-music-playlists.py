import functools


class Solution:
    def numMusicPlaylists(self, N: int, L: int, K: int) -> int:
        @functools.lru_cache(None)
        def dp(i, j):
            if i == 0:
                return int(j == 0)
            res = dp(i - 1, j - 1) * (N - j + 1)
            res += dp(i - 1, j) * max(j - K, 0)
            return res % (10 ** 9 + 7)

        return dp(L, N)
