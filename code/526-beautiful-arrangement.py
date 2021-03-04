class Solution:
    def countArrangement(self, N: int) -> int:
        visited = [False] * (N + 1)
        res = 0

        def dfs(idx):
            nonlocal res
            if idx > N:
                res += 1
            for i in range(1, N + 1):
                if not visited[i] and (idx % i == 0 or i % idx == 0):
                    visited[i] = True
                    dfs(idx + 1)
                    visited[i] = False

        dfs(1)
        return res
