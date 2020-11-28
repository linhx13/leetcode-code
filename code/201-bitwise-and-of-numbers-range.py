class Solution:
    def rangeBitwiseAnd(self, m: int, n: int) -> int:
        k = 0
        while m != n:
            n >>= 1
            m >>= 1
            k += 1
        return m << k


if __name__ == "__main__":
    m = 5
    n = 7
    print(Solution().rangeBitwiseAnd(m, n))
