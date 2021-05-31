from typing import List
import collections


class Solution:
    def suggestedProducts(
        self, products: List[str], searchWord: str
    ) -> List[List[str]]:
        m = collections.defaultdict(list)
        for p in products:
            for i in range(1, len(p) + 1):
                m[p[:i]].append(p)
        for k, v in m.items():
            m[k] = sorted(v)[:3]
        res = []
        for i in range(1, len(searchWord) + 1):
            res.append(m[searchWord[:i]])
        return res
