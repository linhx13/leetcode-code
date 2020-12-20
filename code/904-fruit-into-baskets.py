from typing import List
import collections


class Solution:
    def totalFruit(self, tree: List[int]) -> int:
        i, j = 0, 0
        counter = collections.Counter()
        res = 0
        cnt = 0
        while i < len(tree):
            counter[tree[i]] += 1
            cnt += 1
            while j <= i and len(counter) > 2:
                counter[tree[j]] -= 1
                cnt -= 1
                if counter[tree[j]] == 0:
                    del counter[tree[j]]
                j += 1
            res = max(res, cnt)
            i += 1
        return res


if __name__ == "__main__":
    # tree = [1, 2, 1]
    # tree = [0, 1, 2, 2]
    # tree = [1, 2, 3, 2, 2]
    tree = [3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4]
    # tree = [1, 1, 1, 2, 2, 1]
    print(Solution().totalFruit(tree))
