from typing import List


class Solution:
    def minHeightShelves(self, books: List[List[int]],
                         shelf_width: int) -> int:
        dp = [0 for _ in range(len(books) + 1)]
        for i in range(len(books) - 1, -1, -1):
            dp[i] = dp[i + 1] + books[i][1]
            w, h = books[i]
            for j in range(i + 1, len(books)):
                if w + books[j][0] > shelf_width:
                    break
                w += books[j][0]
                h = max(h, books[j][1])
                dp[i] = min(dp[i], dp[j + 1] + h)
        return dp[0]
