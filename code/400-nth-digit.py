class Solution:
    def findNthDigit(self, n: int) -> int:
        d, c = 1, 9
        while True:
            if n < d * c:
                break
            n -= d * c
            c *= 10
            d += 1
        x = 10 ** (d - 1) + ((n - 1) // d)
        y = (n - 1) % d
        res = (x % (10 ** (d - y))) // (10 ** (d - y - 1))
        return res


if __name__ == "__main__":
    n = 100
    print(Solution().findNthDigit(n))

    c = 0
    num = 1
    while c < n:
        c += len(str(num))
        num += 1
    print(num)
