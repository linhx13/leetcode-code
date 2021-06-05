from typing import List


class Solution:
    def evaluate(self, s: str, knowledge: List[List[str]]) -> str:
        d = {k: v for k, v in knowledge}
        res = []
        cur = ""
        going = False
        for c in s:
            if c == "(":
                going = True
            elif c == ")":
                going = False
                res.append(d.get(cur, "?"))
                cur = ""
            elif going:
                cur += c
            else:
                res.append(c)
        return "".join(res)
