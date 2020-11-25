from typing import List


class Solution:
    def palindromePairs(self, words: List[str]) -> List[List[int]]:
        d = {}
        for i, w in enumerate(words):
            d[w] = i
        res = set()
        for i, w in enumerate(words):
            for j in range(len(w) + 1):
                prefix, suffix = w[:j], w[j:]
                if self.is_pal(prefix):
                    suffix_rev = suffix[::-1]
                    k = d.get(suffix_rev)
                    if k is not None and k != i:
                        res.add((k, i))
                if self.is_pal(suffix):
                    prefix_rev = prefix[::-1]
                    k = d.get(prefix_rev)
                    if k is not None and k != i:
                        res.add((i, k))
        return list(res)

    def is_pal(self, s):
        return s == s[::-1]


if __name__ == "__main__":
    words = ["abcd", "dcba", "lls", "s", "sssll"]
    # words = ["bat", "tab", "cat"]
    # words = ["a", ""]
    print(Solution().palindromePairs(words))
