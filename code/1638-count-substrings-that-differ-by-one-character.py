class Solution:
    def countSubstrings(self, s: str, t: str) -> int:
        m, n = len(s), len(t)

        def helper(i, j):
            res = pre = cur = 0
            for k in range(min(m - i, n - j)):
                cur += 1
                if s[i + k] != t[j + k]:
                    pre, cur = cur, 0
                res += pre
            return res

        return sum(helper(i, 0) for i in range(m)) + sum(
            helper(0, j) for j in range(1, n)
        )
