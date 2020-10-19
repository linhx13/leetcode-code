from typing import List


class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        res = 0
        for l in range(1, len(arr) + 1, 2):
            for i in range(len(arr) - l + 1):
                s = sum(arr[i:i + l])
                res += s
        return res


if __name__ == "__main__":
    arr = [1, 4, 2, 5, 3]
    sol = Solution()
    print(sol.sumOddLengthSubarrays(arr))
