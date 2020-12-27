class Solution:
    def getSmallestString(self, n: int, k: int) -> str:
        res = []
        while k > 0 and n > 0:
            cur = min(k - n, 25)
            res.append(chr(ord("a") + cur))
            k -= cur + 1
            n -= 1
        return "".join(res[::-1])


if __name__ == "__main__":
    # n = 5
    # k = 73
    n = 3
    k = 27
    print(Solution().getSmallestString(n, k))
