from typing import List


class Solution:
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        allowed = set(list(allowed))
        return sum(1 for w in words if not set(list(w)).difference(allowed))
