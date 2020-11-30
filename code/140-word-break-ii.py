from typing import List


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        word_set = set(wordDict)
        len_set = set(len(w) for w in word_set)
        n = len(s)
        memo = {}

        def helper(idx):
            if idx == n:
                return []
            if idx in memo:
                return memo[idx]
            res = []
            for l in len_set:
                if idx + l <= n and s[idx:idx + l] in word_set:
                    if idx + l == n:
                        res.append(s[idx:idx + l])
                    else:
                        for x in helper(idx + l):
                            res.append(s[idx:idx + l] + " " + x)
            memo[idx] = res
            return res

        res = helper(0)
        return res


if __name__ == "__main__":
    # s = "catsanddog"
    # wordDict = ["cat", "cats", "and", "sand", "dog"]
    # s = "pineapplepenapple"
    # wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
    s = "catsandog"
    wordDict = ["cats", "dog", "sand", "and", "cat"]
    print(Solution().wordBreak(s, wordDict))
