class Solution:
    def maxRepeating(self, sequence: str, word: str) -> int:
        res = 0
        while True:
            s = word * (res + 1)
            if sequence.find(s) == -1:
                break
            res += 1
        return res
