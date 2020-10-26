class Solution:
    def new21Game(self, N: int, K: int, W: int) -> float:
        if K == 0 or N >= K + W:
            return 1.0
        dp = [1.0] + [0.0] * N
        res = 0
        wsum = 1.0
        for i in range(1, N + 1):
            dp[i] = wsum / W
            if i < K:
                wsum += dp[i]
            else:
                res += dp[i]
            if i >= W:
                wsum -= dp[i - W]
        return res


if __name__ == '__main__':
    # N = 10
    # K = 1
    # W = 10
    # N = 6
    # K = 1
    # W = 10
    N = 21
    K = 17
    W = 10
    print(Solution().new21Game(N, K, W))
