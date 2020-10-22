from collections import deque


class Solution:
    def magicalString(self, n: int) -> int:
        if n == 0:
            return 0
        queue = deque([2])
        i = 1
        res = 1
        k = 0
        n -= 2
        while n > 0:
            k = queue.popleft()
            if k == 1:
                res += 1
            n -= 1
            if n == 0:
                break
            for _ in range(k):
                queue.append(i)
            if i == 2:
                i = 1
            else:
                i = 2
        return res


if __name__ == "__main__":
    n = 6
    print(Solution().magicalString(n))
