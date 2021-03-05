class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        size = 1
        for i in range(1, n + 1):
            size *= i

        visited = [False] * (n + 1)

        def dfs(x, n, k, size, buf):
            if x == 0:
                return "".join(map(str, buf))
            size //= x

            i = (k + size - 1) // size
            k -= size * (i - 1)
            for j in range(1, n + 1):
                if not visited[j]:
                    i -= 1
                    if i == 0:
                        break
            buf.append(j)
            visited[j] = True
            return dfs(x - 1, n, k, size, buf)

        return dfs(n, n, k, size, [])


if __name__ == "__main__":
    n = 9
    k = 296662
    print(Solution().getPermutation(n, k))
