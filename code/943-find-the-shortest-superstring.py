from typing import List


class Solution:
    def shortestSuperstring(self, A: List[str]) -> str:
        N = len(A)

        overlaps = [[0] * N for _ in range(N)]
        for i, x in enumerate(A):
            for j, y in enumerate(A):
                if i != j:
                    for k in range(min(len(x), len(y)), -1, -1):
                        if x.endswith(y[:k]):
                            overlaps[i][j] = k
                            break

        dp = [[0] * N for _ in range(1 << N)]
        prev = [[None] * N for _ in range(1 << N)]
        for mask in range(1, 1 << N):
            for j in range(N):
                if (mask >> j) & 1:
                    prev_mask = mask ^ (1 << j)
                    if prev_mask == 0:
                        continue
                    for i in range(N):
                        if (prev_mask >> i) & 1:
                            val = dp[prev_mask][i] + overlaps[i][j]
                            if val > dp[mask][j]:
                                dp[mask][j] = val
                                prev[mask][j] = i

        perm = []
        mask = (1 << N) - 1
        i = max(range(N), key=dp[-1].__getitem__)
        while i is not None:
            perm.append(i)
            mask, i = mask ^ (1 << i), prev[mask][i]

        perm = perm[::-1]
        seen = [False] * N
        for x in perm:
            seen[x] = True
        perm.extend([i for i in range(N) if not seen[i]])

        res = [A[perm[0]]]
        for i in range(1, len(perm)):
            overlap = overlaps[perm[i - 1]][perm[i]]
            res.append(A[perm[i]][overlap:])

        return "".join(res)


if __name__ == "__main__":
    A = ["catg", "ctaagt", "gcta", "ttca", "atgcatc"]
    print(Solution().shortestSuperstring(A))
