class Solution:
    def longestNiceSubstring(self, s: str) -> str:
        if not s:
            return ""
        ss = set(s)
        for i, c in enumerate(s):
            if c.swapcase() not in ss:
                s1 = self.longestNiceSubstring(s[:i])
                s2 = self.longestNiceSubstring(s[i + 1 :])
                return max(s1, s2, key=len)
        return s
