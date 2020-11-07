from typing import List


class Solution:
    def canFormArray(self, arr: List[int], pieces: List[List[int]]) -> bool:
        d = {}
        for p in pieces:
            d[p[0]] = p
        idx = 0
        while idx < len(arr):
            if arr[idx] not in d:
                return False
            p = d[arr[idx]]
            for i in range(len(p)):
                if arr[idx] != p[i]:
                    return False
                idx += 1
        return True
