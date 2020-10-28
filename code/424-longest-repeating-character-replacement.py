class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        cnts = {}
        i = 0
        ans = 0
        for j in range(len(s)):
            cnts[s[j]] = cnts.get(s[j], 0) + 1
            if j - i + 1 - max(cnts.values()) > k:
                cnts[s[i]] -= 1
                i += 1
            ans = max(ans, j - i + 1)
        return ans


if __name__ == "__main__":
    s = "AABABBA"
    k = 1
    print(Solution().characterReplacement(s, k))
