from typing import List


class Solution:
    def findAndReplacePattern(
        self, words: List[str], pattern: str
    ) -> List[str]:
        res = []
        for word in words:
            if len(word) != len(pattern):
                continue
            m1 = {}
            m2 = {}
            valid = True
            for i in range(len(pattern)):
                a, b = word[i], pattern[i]
                va = m1.get(a)
                vb = m2.get(b)
                if (va is None and vb is not None) or (
                    va is not None and vb is None
                ):
                    valid = False
                    break
                elif (
                    va is not None
                    and vb is not None
                    and not (va == b and vb == a)
                ):
                    valid = False
                    break
                else:
                    m1[a], m2[b] = b, a
            if valid:
                res.append(word)
        return res


if __name__ == "__main__":
    words = ["abc", "deq", "mee", "aqq", "dkd", "ccc"]
    pattern = "abb"
    print(Solution().findAndReplacePattern(words, pattern))
