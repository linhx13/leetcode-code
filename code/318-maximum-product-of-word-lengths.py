from typing import List


class Solution:
    def maxProduct(self, words: List[str]) -> int:
        bm = [0] * len(words)
        for i, word in enumerate(words):
            x = 0
            for c in word:
                x |= 1 << (ord(c) - ord("a"))
            bm[i] = x

        res = 0
        for i in range(len(words)):
            for j in range(i + 1, len(words)):
                if (bm[i] & bm[j]) == 0:
                    res = max(len(words[i]) * len(words[j]), res)
        return res
