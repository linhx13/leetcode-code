from typing import List
import bisect


class Solution:
    def maxEnvelopes(self, envelopes: List[List[int]]) -> int:
        if not envelopes:
            return 0
        envelopes.sort(key=lambda e: (e[0], -e[1]))
        len_min = [0 for _ in range(len(envelopes))]

        l, r = 0, 0

        for i in range(0, len(envelopes)):
            if r == 0 or len_min[r - 1] < envelopes[i][1]:
                r += 1
                len_min[r - 1] = envelopes[i][1]
            else:
                k = bisect.bisect_left(len_min, envelopes[i][1], l, r)
                len_min[k] = envelopes[i][1]
        return r


if __name__ == "__main__":
    # envelopes = [[5, 4], [6, 4], [6, 7], [2, 3]]
    # envelopes = [[1, 1]]
    envelopes = [
        [15, 8],
        [2, 20],
        [2, 14],
        [4, 17],
        [8, 19],
        [8, 9],
        [5, 7],
        [11, 19],
        [8, 11],
        [13, 11],
        [2, 13],
        [11, 19],
        [8, 11],
        [13, 11],
        [2, 13],
        [11, 19],
        [16, 1],
        [18, 13],
        [14, 17],
        [18, 19],
    ]

    print(Solution().maxEnvelopes(envelopes))
