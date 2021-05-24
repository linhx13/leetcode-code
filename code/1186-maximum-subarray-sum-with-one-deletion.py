from typing import List


class Solution:
    def maximumSum(self, arr: List[int]) -> int:
        n = len(arr)
        if n == 1:
            return arr[0]
        right = [0] * n
        res = right[n - 1] = arr[n - 1]
        for i in range(n - 2, -1, -1):
            right[i] = max(arr[i], right[i + 1] + arr[i])
            res = max(res, right[i])
        left = [0] * n
        left[0] = arr[0]
        for i in range(1, n):
            left[i] = max(arr[i], left[i - 1] + arr[i])
            lv = left[i - 1] if i - 1 >= 0 else 0
            rv = right[i + 1] if i + 1 < n else 0
            res = max(lv + rv, res)
        return res


if __name__ == "__main__":
    arr = [1, -2, 0, 3]
    # arr = [1, -2, -2, 3]
    # arr = [-1, -1, -1, -1]
    print(Solution().maximumSum(arr))
