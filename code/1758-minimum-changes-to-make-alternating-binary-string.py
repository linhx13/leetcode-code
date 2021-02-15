class Solution:
    def minOperations(self, s: str) -> int:
        chs = ["0", "1"]

        def check(idx):
            cnt = 0
            for i in range(len(s)):
                if chs[idx] != s[i]:
                    cnt += 1
                idx = 1 - idx
            return cnt

        return min(check(0), check(1))
