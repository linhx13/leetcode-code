from typing import List


class Solution:
    def decrypt(self, code: List[int], k: int) -> List[int]:
        ans = []
        n = len(code)
        for i in range(n):
            x = 0
            if k > 0:
                x = sum(code[(i + j + n) % n] for j in range(1, k + 1))
            elif k < 0:
                x = sum(code[(i + j + n) % n] for j in range(-1, k - 1, -1))
            else:
                x = 0
            ans.append(x)
        return ans
