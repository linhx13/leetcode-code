import collections


class Solution:
    def minDeletions(self, s: str) -> int:
        counter = collections.Counter(s)
        res = 0
        cnts = set()
        for k, c in counter.most_common():
            while c > 0 and c in cnts:
                c -= 1
                res += 1
            cnts.add(c)
        return res


if __name__ == "__main__":
    s = "bbcebab"
    print(Solution().minDeletions(s))
