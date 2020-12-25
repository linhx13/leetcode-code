from typing import List
from operator import add, sub, mul, truediv


class Solution:
    def judgePoint24(self, A: List[int]) -> bool:
        if not A:
            return False
        if len(A) == 1:
            return abs(A[0] - 24) < 1e-6

        for i in range(len(A)):
            for j in range(len(A)):
                if i != j:
                    B = [A[k] for k in range(len(A)) if i != k and j != k]
                    for op in (truediv, mul, add, sub):
                        if (op is add or op is mul) and j > i:
                            continue
                        if op is not truediv or A[j]:
                            B.append(op(A[i], A[j]))
                            if self.judgePoint24(B):
                                return True
                            B.pop()
        return False
