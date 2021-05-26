class Solution:
    def maxUniqueSplit(self, s: str) -> int:
        n = len(s)
        res = 0
        strs = set()

        def helper(i):
            nonlocal res
            if i >= n:
                res = max(res, len(strs))
                return
            for j in range(i + 1, n + 1):
                w = s[i:j]
                if w not in strs:
                    strs.add(w)
                    helper(j)
                    strs.remove(w)

        helper(0)
        return res
