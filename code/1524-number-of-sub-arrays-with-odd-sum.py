from typing import List


class Solution:
    def numOfSubarrays(self, arr: List[int]) -> int:
        kmod = 10**9 + 7
        prefix = [0]
        for n in arr:
            prefix.append(prefix[-1] + n)
        ans = 0
        odd, even = 0, 0
        for s in prefix:
            if s % 2 == 1:
                ans = (ans + even) % kmod
                odd += 1
            else:
                ans = (ans + odd) % kmod
                even += 1
        return ans


if __name__ == "__main__":
    # arr = [1, 3, 5]
    arr = [100, 100, 99, 99]
    print(Solution().numOfSubarrays(arr))
