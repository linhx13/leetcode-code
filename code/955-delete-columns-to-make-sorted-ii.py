from typing import List
import collections


class Solution:
    def minDeletionSize(self, A: List[str]) -> int:
        res = set()

        def helper(rows, col):
            if col >= len(A[0]):
                return
            if col in res:
                helper(rows, col + 1)
                return
            todos = collections.defaultdict(set)
            for i in range(1, len(rows)):
                cur = A[rows[i]][col]
                prev = A[rows[i - 1]][col]
                if cur < prev:
                    res.add(col)
                    todos["#"].update(rows)
                    break
                elif cur == prev:
                    todos[cur].add(rows[i])
                    todos[cur].add(rows[i - 1])
            for v in todos.values():
                helper(sorted(v), col + 1)

        helper(list(range(len(A))), 0)
        return len(res)


if __name__ == "__main__":
    # A = ["ca", "bb", "ac"]
    # A = ["xc", "yb", "za"]
    # A = ["zyx", "wvu", "tsr"]
    A = ["xga", "xfb", "yfa"]
    # A = [
    #     "bwwdyeyfhc",
    #     "bchpphbtkh",
    #     "hmpudwfkpw",
    #     "lqeoyqkqwe",
    #     "riobghmpaa",
    #     "stbheblgao",
    #     "snlaewujlc",
    #     "tqlzolljas",
    #     "twdkexzvfx",
    #     "wacnnhjdis",
    # ]

    print(Solution().minDeletionSize(A))
