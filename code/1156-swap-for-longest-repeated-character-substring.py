class Solution:
    def maxRepOpt1(self, text: str) -> int:
        freqs = {}
        last, cnt = None, 0
        chars = []
        for i in range(len(text)):
            freqs[text[i]] = freqs.get(text[i], 0) + 1
            if text[i] != last and cnt > 0:
                chars.append((last, cnt))
                cnt = 0
            last, cnt = text[i], cnt + 1
        if cnt > 0:
            chars.append((last, cnt))
        res = 1
        for i in range(len(chars)):
            if chars[i][1] == 1:
                if (i - 1 >= 0 and i + 1 < len(chars) \
                   and chars[i - 1][0] == chars[i + 1][0]):
                    c = chars[i - 1][1] + chars[i + 1][1]
                    if freqs[chars[i - 1][0]] > c:
                        c += 1
                    res = max(res, c)
            else:
                c = chars[i][1]
                if freqs[chars[i][0]] > chars[i][1]:
                    c += 1
                res = max(res, c)
        return res
