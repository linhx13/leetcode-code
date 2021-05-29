from typing import List


class Solution:
    def rotateTheBox(self, box: List[List[str]]) -> List[List[str]]:
        m, n = len(box), len(box[0])
        res = [["" for _ in range(m)] for _ in range(n)]
        for r in range(m):
            for c in range(n):
                res[c][r] = box[m - 1 - r][c]
        for c in range(m):
            i, j = n - 1, n - 1
            while i >= 0 and j >= 0:
                if res[i][c] == "*":
                    i, j = i - 1, i - 1
                    continue
                elif res[i][c] == "#":
                    res[i][c] = "."
                    res[j][c] = "#"
                    j -= 1
                i -= 1
        return res
