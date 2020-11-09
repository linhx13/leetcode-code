from collections import Counter


class Solution:
    def originalDigits(self, s: str) -> str:
        counter = Counter(s)
        digits = [
            "zero", "one", "two", "three", "four", "five", "six", "seven",
            "eight", "nine"
        ]
        freq = [0] * 10
        for x, i in ("z", 0), ("w",
                               2), ("u",
                                    4), ("x",
                                         6), ("g",
                                              8), ("s",
                                                   7), ("f",
                                                        5), ("o",
                                                             1), ("h",
                                                                  3), ("i", 9):
            freq[i] += counter[x]
            counter -= Counter(digits[i] * counter[x])
        return "".join(str(i) * x for i, x in enumerate(freq))
