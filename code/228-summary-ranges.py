from typing import List


class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        start, end = None, None
        res = []
        for n in nums:
            if start is None:
                start, end = n, n
                continue
            if n == end + 1:
                end = n
            else:
                if start == end:
                    res.append(str(start))
                else:
                    res.append("%s->%s" % (start, end))
                start, end = n, n
        if start is not None:
            if start == end:
                res.append(str(start))
            else:
                res.append("%s->%s" % (start, end))
        return res
