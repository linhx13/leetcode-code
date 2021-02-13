from typing import List


class Solution:
    def maxSatisfied(
        self, customers: List[int], grumpy: List[int], X: int
    ) -> int:
        a, b, c = 0, 0, 0
        res = 0
        total = 0
        m = 0
        for i in range(len(customers)):
            if grumpy[i] == 0:
                total += customers[i]
                a += customers[i]
            b += customers[i]
            if c < X:
                c += 1
            if c == X:
                if i >= X:
                    if grumpy[i - X] == 0:
                        a -= customers[i - X]
                    b -= customers[i - X]
                m = max(m, -a + b)
        return m + total


if __name__ == "__main__":
    customers = [4, 10, 10]
    grumpy = [1, 1, 0]
    X = 2
    print(Solution().maxSatisfied(customers, grumpy, X))
