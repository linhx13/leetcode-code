class Solution:
    def secondHighest(self, s: str) -> int:
        first, second = -1, -1
        for c in s:
            if not c.isnumeric():
                continue
            x = int(c)
            if first == -1:
                first = x
            else:
                if x > first:
                    first, second = x, first
                elif x == first:
                    continue
                elif x > second:
                    second = x
        return second
