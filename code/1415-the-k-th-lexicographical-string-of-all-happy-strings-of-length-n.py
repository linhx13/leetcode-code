class Solution:
    def getHappyString(self, n: int, k: int) -> str:
        chars = ["a", "b", "c"]

        def helper(n, i, buf):
            nonlocal k
            if n == 0:
                k -= 1
                if k == 0:
                    return "".join(chars[j] for j in buf)
                else:
                    return ""
            for j in range(0, 3):
                if j != i:
                    buf.append(j)
                    res = helper(n - 1, j, buf)
                    if res:
                        return res
                    buf.pop()
            return ""

        return helper(n, -1, [])


if __name__ == "__main__":
    # n = 1
    # k = 3
    n = 10
    k = 100
    print(Solution().getHappyString(n, k))
