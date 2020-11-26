from typing import List


class Solution:
    def pancakeSort(self, arr: List[int]) -> List[int]:
        def flip(sublist, k):
            i = 0
            while i < k // 2:
                sublist[i], sublist[k - i - 1] = sublist[k - i - 1], sublist[i]
                i += 1

        ans = []
        value_to_sort = len(arr)
        while value_to_sort > 0:
            idx = arr.index(value_to_sort)
            if idx != value_to_sort - 1:
                if idx != 0:
                    ans.append(idx + 1)
                    flip(arr, idx + 1)
                ans.append(value_to_sort)
                flip(arr, value_to_sort)
            value_to_sort -= 1
        return ans
