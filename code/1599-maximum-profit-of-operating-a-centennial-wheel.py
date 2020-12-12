from typing import List


class Solution:
    def minOperationsMaxProfit(
        self, customers: List[int], boardingCost: int, runningCost: int
    ) -> int:
        res = -1
        waiting = 0
        total_profit = 0
        max_profit = 0
        for i, x in enumerate(customers):
            waiting += x
            cur = min(4, waiting)
            waiting -= cur
            total_profit += cur * boardingCost - runningCost
            if max_profit < total_profit:
                res = i + 1
                max_profit = total_profit
        q, r = divmod(waiting, 4)
        if 4 * boardingCost > runningCost:
            res += q
        if r * boardingCost > runningCost:
            res += 1
        return res
