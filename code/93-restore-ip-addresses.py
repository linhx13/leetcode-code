from typing import List


class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        res = []
        self.helper(s, 0, [], res)
        return res

    def helper(self, s, idx, buf, res):
        if idx == len(s) and len(buf) == 4:
            res.append(".".join(buf))
            return
        if idx >= len(s) or len(buf) >= 4:
            return
        for i in range(idx, len(s)):
            n = int(s[idx:i + 1])
            if n > 255:
                break
            cur = s[idx:i + 1]
            if n == 0 and len(cur) > 1:
                break
            if n != 0 and cur[0] == "0":
                break
            buf.append(cur)
            self.helper(s, i + 1, buf, res)
            buf.pop()


if __name__ == "__main__":
    # s = "25525511135"
    # s = "0000"
    s = "1111"
    # s = "010010"
    # s = "101023"
    print(Solution().restoreIpAddresses(s))
