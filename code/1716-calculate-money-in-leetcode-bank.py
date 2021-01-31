class Solution:
    def totalMoney(self, n: int) -> int:
        k = n // 7
        res = sum(i for i in range(1, 8)) * k + k * max(0, k - 1) // 2 * 7
        res += sum((k + i) for i in range(1, n % 7 + 1))
        return res


if __name__ == "__main__":
    n = 4
    print(Solution().totalMoney(n))
