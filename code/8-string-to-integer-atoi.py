class Solution:
    def myAtoi(self, s: str) -> int:
        num = ""
        sign = ""
        for ch in s:
            if ch == ' ':
                if len(sign) != 0:
                    break
            elif (ch == '+' or ch == '-') and len(sign) == 0:
                sign += ch
            elif ord('0') <= ord(ch) <= ord('9'):
                num += ch
                if len(sign) == 0:
                    sign = "+"
            else:
                break
        if num == "":
            return 0
        if sign == "+":
            return min(int(num), 2**31 - 1)
        else:
            return max(-int(num), -2**31)
