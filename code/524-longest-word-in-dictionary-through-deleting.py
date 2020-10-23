from typing import List


class Solution:
    def findLongestWord(self, s: str, d: List[str]) -> str:
        res = ""
        for t in d:
            if self.check(s, t):
                if len(t) > len(res) or (len(t) == len(res) and t < res):
                    res = t
        return res

    def check(self, s, t):
        j = 0
        for i in range(len(s)):
            if s[i] == t[j]:
                j += 1
                if j >= len(t):
                    break
        return j == len(t)
