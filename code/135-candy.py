from typing import List


class Solution:
    def candy(self, ratings: List[int]) -> int:
        cnts = [1] * len(ratings)
        for i in range(1, len(ratings)):
            if ratings[i] > ratings[i - 1]:
                cnts[i] = cnts[i - 1] + 1
        res = cnts[-1]
        for i in range(len(ratings) - 2, -1, -1):
            if ratings[i] > ratings[i + 1]:
                cnts[i] = max(cnts[i], cnts[i + 1] + 1)
            res += cnts[i]
        return res


if __name__ == "__main__":
    ratings = [1, 0, 2]
    # ratings = [1, 1, 1]
    print(Solution().candy(ratings))
