from typing import List


class Solution:
    def minOperations(self, logs: List[str]) -> int:
        depth = 0
        for x in logs:
            if x.startswith(".."):
                depth = max(0, depth - 1)
            elif x.startswith("."):
                continue
            else:
                depth += 1
        return depth
