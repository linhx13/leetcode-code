class Solution:
    def nearestPalindromic(self, n: str) -> str:
        def mirroring(s):
            x = s[0:len(s) // 2]
            if len(s) % 2 == 1:
                y = s[len(s) // 2]
            else:
                y = ""
            return x + y + x[::-1]

        if n == "1":
            return "0"

        max_int = 99999999999999999999
        a = mirroring(n)
        diff1 = max_int
        diff1 = abs(int(n) - int(a))
        if diff1 == 0:
            diff1 = max_int

        s = list(n)
        i = (len(s) - 1) // 2
        while i >= 0 and s[i] == '0':
            s[i] = '9'
            i -= 1
        if i == 0 and s[i] == '1':
            s = s[1:]
            mid = (len(s) - 1) // 2
            s[mid] = '9'
        else:
            s[i] = str(int(s[i]) - 1)
        b = mirroring("".join(s))
        diff2 = abs(int(n) - int(b))

        s = list(n)
        i = (len(s) - 1) // 2
        while i >= 0 and s[i] == '9':
            s[i] = '0'
            i -= 1
        if i < 0:
            s.insert(0, '1')
        else:
            s[i] = str(int(s[i]) + 1)
        c = mirroring("".join(s))
        diff3 = abs(int(n) - int(c))

        if diff2 <= diff1 and diff2 <= diff3:
            return b
        if diff1 <= diff3 and diff1 <= diff2:
            return a
        else:
            return c
