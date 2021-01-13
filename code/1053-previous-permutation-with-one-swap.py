from typing import List


class Solution:
    def prevPermOpt1(self, arr: List[int]) -> List[int]:
        m = float("inf")
        i = len(arr) - 1
        while i >= 0:
            if arr[i] > m:
                break
            m = min(m, arr[i])
            i -= 1
        if i < 0:
            return arr
        k = -1
        for j in range(i + 1, len(arr)):
            if arr[j] < arr[i]:
                if k == -1 or arr[j] > arr[k]:
                    k = j
        arr[i], arr[k] = arr[k], arr[i]
        return arr


if __name__ == "__main__":
    # arr = [3, 2, 1]
    # arr = [1, 1, 5]
    arr = [1, 9, 4, 7, 6]
    # arr = [3, 1, 1, 3]
    print(Solution().prevPermOpt1(arr))
