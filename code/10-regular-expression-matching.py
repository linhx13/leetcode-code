import functools


class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        @functools.lru_cache(None)
        def match(i, j):
            if j == len(p):
                return i == len(s)
            else:
                first_match = i < len(s) and p[j] in [s[i], "."]
                if j + 1 < len(p) and p[j + 1] == "*":
                    return match(i, j + 2) or (first_match and match(i + 1, j))
                else:
                    return first_match and match(i + 1, j + 1)

        return match(0, 0)


if __name__ == "__main__":
    s = "aa"
    p = "a"
    # s = "aa"
    # p = "a*"
    s = "ab"
    p = ".*"
    s = "aab"
    p = "c*a*b"
    # s = "mississippi"
    # p = "mis*is*p*."
    # s = "a"
    # p = "ab*"
    s = "a"
    p = ".*..a*"
    print(Solution().isMatch(s, p))
