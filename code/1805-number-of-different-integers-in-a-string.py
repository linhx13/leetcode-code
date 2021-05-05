class Solution:
    def numDifferentIntegers(self, word: str) -> int:
        s = set()
        start = -1
        for i, c in enumerate(word):
            if "0" <= c <= "9":
                if start == -1:
                    start = i
            else:
                if start != -1:
                    s.add(int(word[start:i]))
                    start = -1
        if start != -1:
            s.add(int(word[start:]))
        return len(s)
