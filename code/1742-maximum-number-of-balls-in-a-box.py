from collections import defaultdict


class Solution:
    def countBalls(self, lowLimit: int, highLimit: int) -> int:
        freq = defaultdict(int)
        for x in range(lowLimit, highLimit + 1):
            k = sum(int(i) for i in str(x))
            freq[k] += 1
        return max(freq.values())
