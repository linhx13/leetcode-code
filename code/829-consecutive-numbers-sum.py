import math


class Solution:
    def consecutiveNumbersSum(self, N: int) -> int:
        # res = 0
        # n = 1
        # while True:
        #     avg = math.ceil(N / n)
        #     minn = avg - n // 2
        #     if minn <= 0:
        #         break
        #     if (minn + minn + n - 1) * n / 2 == N:
        #         res += 1
        #     elif (minn + 1 + minn + 1 + n - 1) * n / 2 == N:
        #         res += 1
        #     maxn = minn + n - 1
        #     n += 1
        # return res
        res = 1
        for k in range(2, int(math.sqrt(2 * N))):
            if (N - k * (k - 1) // 2) % k == 0:
                res += 1
        return res


if __name__ == "__main__":
    N = 15
    print(Solution().consecutiveNumbersSum(N))
