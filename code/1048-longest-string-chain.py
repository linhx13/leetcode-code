from typing import List
from collections import defaultdict


class Solution:
    def longestStrChain(self, words: List[str]) -> int:
        words.sort(key=lambda w: -len(w))
        m = set(words)
        dp = defaultdict(int)
        for word in words:
            dp[word] = max(0, dp[word])
            for k in range(len(word)):
                key = word[:k] + word[k + 1 :]
                if key not in m:
                    continue
                dp[key] = max(dp[key], dp[word] + 1)

        return max(dp.values()) + 1
