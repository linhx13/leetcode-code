from typing import List


class Solution:
    def shoppingOffers(self, price: List[int], special: List[List[int]],
                       needs: List[int]) -> int:
        memo = {}
        return self.helper(price, special, needs, memo)

    def helper(self, price, special, needs, memo):
        key = tuple(needs)
        if key in memo:
            return memo[key]
        res = sum(needs[i] * price[i] for i in range(len(needs)))
        for s in special:
            clone = needs[:]
            for i in range(len(needs)):
                diff = clone[i] - s[i]
                if diff < 0:
                    break
                clone[i] = diff
            else:
                x = s[-1] + self.helper(price, special, clone, memo)
                res = min(res, x)
        memo[key] = res
        return res
