from typing import List


class Solution:
    def averageWaitingTime(self, customers: List[List[int]]) -> float:
        total = 0
        cur = 0
        for c in customers:
            cur = (c[0] + c[1]) if cur < c[0] else (cur + c[1])
            total += cur - c[0]
        return total / len(customers)


if __name__ == "__main__":
    customers = [[5, 2], [5, 4], [10, 3], [20, 1]]
    # customers = [[1, 2], [2, 5], [4, 3]]
    print(Solution().averageWaitingTime(customers))
