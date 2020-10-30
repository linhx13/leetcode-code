import math


class Solution:
    def primePalindrome(self, N: int) -> int:
        if N <= 2:
            return 2
        if N == 3:
            return 3
        if N <= 5:
            return 5
        if N <= 7:
            return 7
        if N <= 11:
            return 11
        l = len(str(N))
        while True:
            for n in self.gen_pals(l):
                if n >= N and self.is_prime(n):
                    return n
            l += 1

    def gen_pals(self, l):
        ll = l // 2
        for i in range(int(10**(ll - 1)), 10**ll):
            s = str(i)
            for j in range(10):
                yield int(s + str(j) + s[::-1])

    def is_prime(self, n):
        if n <= 3:
            return n > 1
        if n % 6 != 1 and n % 6 != 5:
            return False
        sqrt = int(math.sqrt(n))
        for i in range(5, sqrt + 1, 6):
            if n % i == 0 or n % (i + 2) == 0:
                return False
        return True


if __name__ == "__main__":
    N = 6
    print(Solution().primePalindrome(N))
