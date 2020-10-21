from typing import List


class Solution:
    def findLUSlength(self, strs: List[str]) -> int:
        strs.sort(key=len, reverse=True)
        for i in range(len(strs)):
            if not any(
                    self.check(strs[i], strs[j])
                    for j in range(len(strs)) if i != j):
                return len(strs[i])
        return -1

    def check(self, s, t):
        l = iter(t)
        return all(c in l for c in s)


if __name__ == "__main__":
    strs = ["aabbcc", "aabbcc", "cb"]
    sol = Solution()
    print(sol.findLUSlength(strs))
