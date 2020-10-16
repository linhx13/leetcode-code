from typing import List


class Solution:
    def lexicalOrder(self, n: int) -> List[int]:
        arr = [str(i) for i in range(1, n + 1)]
        arr.sort()
        return [int(x) for x in arr]
