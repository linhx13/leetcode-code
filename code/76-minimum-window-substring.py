import collections


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        diff = collections.Counter()
        for c in t:
            diff[c] -= 1
        t = set(list(t))
        i, j = 0, 0
        res = ""
        less = len(diff)
        while i < len(s):
            if j < i and s[j] in t:
                diff[s[j]] -= 1
                if diff[s[j]] == 0:
                    del diff[s[i]]
                elif diff[s[j]] < 0:
                    less += 1
                j += 1
            while i < len(s) and less > 0:
                if s[i] in t:
                    diff[s[i]] += 1
                    if diff[s[i]] == 0:
                        del diff[s[i]]
                        less -= 1
                i += 1
            while j < i:
                if s[j] in t:
                    if s[j] not in diff:
                        break
                    if s[j] in diff and diff[s[j]] > 0:
                        diff[s[j]] -= 1
                        if diff[s[j]] == 0:
                            del diff[s[j]]
                j += 1
            if less == 0 and j < i:
                window = s[j:i]
                if (not res) or len(window) < len(res):
                    res = window
        return res


if __name__ == "__main__":
    # s = "ADOBECODEBANC"
    # t = "ABC"
    s = "a"
    t = "a"
    # s = "aaaaaaaaaaaabbbbbcdd"
    # t = "abcdd"
    print(Solution().minWindow(s, t))
