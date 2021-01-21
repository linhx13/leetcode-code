class Solution:
    def numSteps(self, s: str) -> int:
        res = 0
        while len(s) > 1:
            if s[-1] == "0":
                s = s[: len(s) - 1]
            else:
                idx = s.rfind("0")
                s = (
                    ("1" + "0" * len(s))
                    if idx == -1
                    else (s[:idx] + "1" + "0" * (len(s) - idx - 1))
                )
            res += 1
        return res
