class Solution:
    def isInterleave(self, s1: str, s2: str, s3: str) -> bool:
        if len(s1) + len(s2) != len(s3):
            return False
        memo = {}
        for i in range(len(s1)):
            for j in range(len(s2)):
                memo[(i, j)] = -1
        return self.helper(s1, s2, s3, 0, 0, 0, memo)

    def helper(self, s1, s2, s3, i, j, k, memo):
        if i == len(s1):
            return s2[j:] == s3[k:]
        if j == len(s2):
            return s1[i:] == s3[k:]
        if memo[(i, j)] >= 0:
            return memo[(i, j)] == 1
        res = False
        if (
            s3[k] == s1[i] and self.helper(s1, s2, s3, i + 1, j, k + 1, memo)
        ) or (
            s3[k] == s2[j] and self.helper(s1, s2, s3, i, j + 1, k + 1, memo)
        ):
            res = True
        memo[(i, j)] = int(res)
        return res
