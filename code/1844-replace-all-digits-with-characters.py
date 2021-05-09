class Solution:
    def replaceDigits(self, s: str) -> str:
        s = list(s)

        def shift(i):
            if i & 1 == 0:
                return s[i]
            else:
                c, x = s[i - 1], int(s[i])
                return chr(ord(c) + x)

        return "".join(map(shift, range(len(s))))
