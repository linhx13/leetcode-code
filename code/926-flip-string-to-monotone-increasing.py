class Solution:
    def minFlipsMonoIncr(self, S: str) -> int:
        sums = [0]
        for x in S:
            sums.append(sums[-1] + int(x))
        return min(sums[i] + len(S) - i - (sums[-1] - sums[i])
                   for i in range(len(sums)))
