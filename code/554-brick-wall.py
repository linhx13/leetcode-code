from typing import List
from collections import Counter


class Solution:
    def leastBricks(self, wall: List[List[int]]) -> int:
        counter = Counter()
        for row in wall:
            total = 0
            for i, x in enumerate(row):
                total += x
                if i != len(row) - 1:
                    counter[total] += 1
        if not counter:
            return len(wall)
        else:
            return len(wall) - counter.most_common()[0][1]


if __name__ == "__main__":
    # wall = [[1, 2, 2, 1], [3, 1, 2], [1, 3, 2], [2, 4], [3, 1, 2],
    #         [1, 3, 1, 1]]
    wall = [[1], [1], [1]]
    print(Solution().leastBricks(wall))
