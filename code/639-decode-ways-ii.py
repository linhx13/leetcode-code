import functools


class Solution:
    def numDecodings(self, s: str) -> int:
        MOD = 1000000007

        @functools.lru_cache(None)
        def dfs(i):
            if i < 0:
                return 1
            if s[i] == "*":
                res = 9 * dfs(i - 1)
                if i > 0 and s[i - 1] == "1":
                    res = (res + 9 * dfs(i - 2)) % MOD
                elif i > 0 and s[i - 1] == "2":
                    res = (res + 6 * dfs(i - 2)) % MOD
                elif i > 0 and s[i - 1] == "*":
                    res = (res + 15 * dfs(i - 2)) % MOD
                return res
            res = dfs(i - 1) if s[i] != "0" else 0
            if i > 0 and s[i - 1] == "1":
                res = (res + dfs(i - 2)) % MOD
            elif i > 0 and s[i - 1] == "2" and s[i] <= "6":
                res = (res + dfs(i - 2)) % MOD
            elif i > 0 and s[i - 1] == "*":
                res = (res + (2 if s[i] <= "6" else 1) * dfs(i - 2)) % MOD
            return res

        return dfs(len(s) - 1) % MOD
