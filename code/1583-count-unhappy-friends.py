from typing import List


class Solution:
    def unhappyFriends(self, n: int, preferences: List[List[int]],
                       pairs: List[List[int]]) -> int:
        rank = [[0 for _ in range(n)] for _ in range(n)]
        for i, p in enumerate(preferences):
            for k, j in enumerate(p):
                rank[i][j] = n - k

        def check(x, y, u, v):
            if rank[x][u] > rank[x][y] and rank[u][x] > rank[u][v]:
                return True
            else:
                return False

        ans = set()
        for i in range(len(pairs)):
            for j in range(len(pairs)):
                if i == j:
                    continue
                if check(pairs[i][0], pairs[i][1], pairs[j][0], pairs[j][1]) \
                   or check(pairs[i][0], pairs[i][1], pairs[j][1], pairs[j][0]):
                    ans.add(pairs[i][0])
                if check(pairs[i][1], pairs[i][0], pairs[j][0], pairs[j][1]) \
                   or (check(pairs[i][1], pairs[i][0], pairs[j][1], pairs[j][0])):
                    ans.add(pairs[i][1])
        return len(ans)


if __name__ == "__main__":
    n = 6
    # preferences = [[1, 2, 3], [3, 2, 0], [3, 1, 0], [1, 2, 0]]
    # pairs = [[0, 1], [2, 3]]
    # preferences = [[1, 3, 2], [2, 3, 0], [1, 3, 0], [0, 2, 1]]
    # pairs = [[1, 3], [0, 2]]
    preferences = [[1, 4, 3, 2, 5], [0, 5, 4, 3, 2], [3, 0, 1, 5, 4],
                   [2, 1, 4, 0, 5], [2, 1, 0, 3, 5], [3, 4, 2, 0, 1]]
    pairs = [[3, 1], [2, 0], [5, 4]]
    print(Solution().unhappyFriends(n, preferences, pairs))
