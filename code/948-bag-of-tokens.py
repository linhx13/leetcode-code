from typing import List
from collections import deque


class Solution:
    def bagOfTokensScore(self, tokens: List[int], P: int) -> int:
        tokens = deque(sorted(tokens))
        points = 0
        res = 0
        while tokens and P >= tokens[0]:
            while tokens and P >= tokens[0]:
                P -= tokens.popleft()
                points += 1
            res = max(res, points)
            while tokens and points and P < tokens[0]:
                P += tokens.pop()
                points -= 1
        return res


if __name__ == "__main__":
    tokens = [100, 200, 300, 400]
    P = 200
    sol = Solution()
    print(sol.bagOfTokensScore(tokens, P))
