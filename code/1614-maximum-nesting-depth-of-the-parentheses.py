class Solution:
    def maxDepth(self, s: str) -> int:
        res = 0
        l, r = 0, 0
        for c in s:
            if c == '(':
                l += 1
            elif c == ')':
                r += 1
            res = max(res, l - r)
        return res
