class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        res = [[] for _ in range(numRows)]
        i = 0
        for _ in range(len(s) // (2 * numRows - 2)):
            for r in range(numRows):
                res[r].append(s[i])
                i += 1
            for r in range(numRows - 2):
                res[numRows - 2 - r].append(s[i])
                i += 1
        for r in range(numRows):
            if i >= len(s):
                break
            res[r].append(s[i])
            i += 1
        for r in range(numRows - 2):
            if i >= len(s):
                break
            res[numRows - 2 - r].append(s[i])
            i += 1
        return "".join("".join(row) for row in res)


if __name__ == "__main__":
    # s = "PAYPALISHIRING"
    # numRows = 3
    # t = "PAHNAPLSIIGYIR"
    s = "PAYPALISHIRING"
    numRows = 4
    t = "PINALSIGYAHRPI"
    assert Solution().convert(s, numRows) == t
