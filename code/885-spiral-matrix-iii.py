from typing import List


class Solution:
    def spiralMatrixIII(
        self, R: int, C: int, r0: int, c0: int
    ) -> List[List[int]]:
        res = []
        total = R * C
        k = 0
        r, c = r0, c0
        res.append((r, c))
        while len(res) < total:
            for _ in range(k):
                r -= 1
                if 0 <= r < R and 0 <= c < C:
                    res.append((r, c))
            for _ in range(k + 1):
                c += 1
                if 0 <= r < R and 0 <= c < C:
                    res.append((r, c))
            for _ in range(k + 1):
                r += 1
                if 0 <= r < R and 0 <= c < C:
                    res.append((r, c))
            for _ in range(k + 1):
                c -= 1
                if 0 <= r < R and 0 <= c < C:
                    res.append((r, c))
            c -= 1
            if 0 <= r < R and 0 <= c < C:
                res.append((r, c))
            k += 2
        return res


if __name__ == "__main__":
    R = 1
    C = 4
    r0 = 0
    c0 = 0
    print(Solution().spiralMatrixIII(R, C, r0, c0))
