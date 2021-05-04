class Solution:
    def checkOnesSegment(self, s: str) -> bool:
        zero = False
        for c in s:
            if c == "0":
                zero = True
            elif c == "1" and zero:
                return False
        return True
