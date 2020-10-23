class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        d = {}
        for i, c in enumerate(s):
            if c not in d:
                d[c] = i
        res = -1
        for i in range(len(s) - 1, -1, -1):
            j = d[s[i]]
            if j != i:
                res = max(res, i - j - 1)
        return res
