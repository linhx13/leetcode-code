from typing import List


class TreeAncestor:
    def __init__(self, n: int, parent: List[int]):
        self.dp = [[-1 for _ in range(n)] for _ in range(20)]
        for i in range(n):
            self.dp[0][i] = parent[i]
        for i in range(1, 20):
            for node in range(n):
                p = self.dp[i - 1][node]
                if p != -1:
                    self.dp[i][node] = self.dp[i - 1][p]

    def getKthAncestor(self, node: int, k: int) -> int:
        for i in range(20):
            if k & (1 << i):
                node = self.dp[i][node]
                if node == -1:
                    return -1
        return node
