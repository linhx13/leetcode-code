from typing import List


class Solution:
    def grayCode(self, n: int) -> List[int]:
        memo = {}

        def helper(k):
            if k == 0:
                return [0]
            if k in memo:
                return memo[k]
            x = helper(k - 1)
            y = [((1 << (k - 1)) | i) for i in x[::-1]]
            memo[k] = x + y
            return memo[k]

        return helper(n)


if __name__ == "__main__":
    n = 3
    print(Solution().grayCode(n))
