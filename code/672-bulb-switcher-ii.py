import itertools


class Solution:
    def flipLights(self, n: int, m: int) -> int:
        res = set()
        for c in itertools.product((0, 1), repeat=4):
            if sum(c) % 2 == m % 2 and sum(c) <= m:
                A = []
                for i in range(min(n, 3)):
                    light = 1
                    light ^= c[0]
                    light ^= c[1] and i % 2 == 1
                    light ^= c[2] and i % 2 == 0
                    light ^= c[3] and i % 3 == 0
                    A.append(light)
                res.add(tuple(A))
        return len(res)
