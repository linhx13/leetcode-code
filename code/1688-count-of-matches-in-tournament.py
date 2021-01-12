class Solution:
    def numberOfMatches(self, n: int) -> int:
        res = 0
        while n > 1:
            if (n & 1) == 1:
                res += ((n - 1) >> 1) + 1
                n = (n - 1) >> 1
            else:
                res += n >> 1
                n = n >> 1
        return res


if __name__ == "__main__":
    n = 14
    print(Solution().numberOfMatches(n))
