from typing import List


class Solution:
    def reconstructMatrix(
        self, upper: int, lower: int, colsum: List[int]
    ) -> List[List[int]]:
        n = len(colsum)
        res = [[0] * n for _ in range(2)]
        for i, s in enumerate(colsum):
            if s == 2 or (s == 1 and lower < upper):
                res[0][i] = 1
            if s == 2 or (s == 1 and not res[0][i]):
                res[1][i] = 1
            upper -= res[0][i]
            lower -= res[1][i]
        if lower == 0 and upper == 0:
            return res
        else:
            return []


if __name__ == "__main__":
    # upper = 2
    # lower = 3
    # colsum = [2, 2, 1, 1]
    upper = 5
    lower = 5
    colsum = [2, 1, 2, 0, 1, 0, 1, 2, 0, 1]
    # upper = 1
    # lower = 4
    # colsum = [2, 1, 2, 0, 0, 2]
    print(Solution().reconstructMatrix(upper, lower, colsum))
