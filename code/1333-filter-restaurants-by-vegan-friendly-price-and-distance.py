from typing import List


class Solution:
    def filterRestaurants(self, restaurants: List[List[int]],
                          veganFriendly: int, maxPrice: int,
                          maxDistance: int) -> List[int]:
        items = []
        for x in restaurants:
            if x[3] > maxPrice or x[4] > maxDistance:
                continue
            if veganFriendly and not x[2]:
                continue
            items.append(x)
        items.sort(key=lambda x: (-x[1], -x[0]))
        return [x[0] for x in items]
