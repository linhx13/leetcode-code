class Solution:
    def reorderedPowerOf2(self, N: int) -> bool:
        s = set()
        p = 1
        while p <= 10 ** 9:
            s.add(tuple(sorted(str(p))))
            p <<= 1
        n = tuple(sorted(str(N)))
        return n in s


if __name__ == "__main__":
    N = 46
    print(Solution().reorderedPowerOf2(N))
