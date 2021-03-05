from typing import List


class Solution:
    def countMatches(
        self, items: List[List[str]], ruleKey: str, ruleValue: str
    ) -> int:
        key2idx = {"type": 0, "color": 1, "name": 2}
        idx = key2idx[ruleKey]
        return sum(1 for item in items if item[idx] == ruleValue)
