from typing import List


class Solution:
    def beautifulArray(self, N: int) -> List[int]:
        if N == 1:
            return [1]
        even = self.beautifulArray(N // 2)
        odd = self.beautifulArray(N - N // 2)
        res = []
        for i in even:
            res.append(i * 2)
        for i in odd:
            res.append(i * 2 - 1)
        return res


if __name__ == "__main__":
    N = 4
    sol = Solution()
    print(sol.beautifulArray(N))
