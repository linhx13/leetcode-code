class Solution:
    def minCut(self, s: str) -> int:
        N = len(s)
        cut = [0] * N
        pal = [[False] * N for _ in range(N)]

        for i in range(N):
            minn = i
            for j in range(i + 1):
                if s[i] == s[j] and (j + 1 > i - 1 or pal[j + 1][i - 1]):
                    pal[j][i] = True
                    minn = 0 if j == 0 else min(minn, cut[j - 1] + 1)
            cut[i] = minn
        return cut[N - 1]
