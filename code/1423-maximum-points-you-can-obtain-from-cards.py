from typing import List


class Solution:
    def maxScore(self, cardPoints: List[int], k: int) -> int:
        n = len(cardPoints)
        total_score = sum(cardPoints)
        window_size = n - k
        window_score = sum(cardPoints[:window_size])
        res = total_score - window_score
        for i in range(window_size, n):
            window_score += cardPoints[i] - cardPoints[i - window_size]
            res = max(res, total_score - window_score)
        return res


if __name__ == "__main__":
    cardPoints = [1, 79, 80, 1, 1, 1, 200, 1]
    k = 3
    print(Solution().maxScore(cardPoints, k))
