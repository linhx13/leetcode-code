class Solution:
    def checkZeroOnes(self, s: str) -> bool:
        ones, zeros = 0, 0
        last = None
        cur = 0
        for c in s:
            if c != last:
                if last == "1":
                    ones = max(ones, cur)
                elif last == "0":
                    zeros = max(zeros, cur)
                cur = 0
            cur += 1
            last = c
        if last == "1":
            ones = max(ones, cur)
        elif last == "0":
            zeros = max(zeros, cur)
        return ones > zeros
