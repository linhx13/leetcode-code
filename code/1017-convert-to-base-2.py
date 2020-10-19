class Solution:
    def baseNeg2(self, N: int) -> str:
        q = N // 2
        r = N % 2
        res = [str(r)]
        k = 1
        while q > 0:
            d = 2 * (-1)**k
            r = q % d
            q //= d
            q = abs(q)
            res.append(str(abs(r)))
            k = 1 - k
        res.reverse()
        return "".join(res)


if __name__ == "__main__":
    N = 2
    sol = Solution()
    print(sol.baseNeg2(N))
