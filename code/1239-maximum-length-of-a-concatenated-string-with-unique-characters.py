from typing import List


class Solution:
    def maxLength(self, arr: List[str]) -> int:
        arr = list(filter(lambda x: len(x) == len(set(x)), arr))
        if not arr:
            return 0
        res = [0]
        self.helper(arr, 0, set(), res)
        return res[0]

    def helper(self, arr, idx, chars, res):
        if idx == len(arr):
            res[0] = max(res[0], len(chars))
            return
        self.helper(arr, idx + 1, chars, res)
        tmp = set(arr[idx])
        if not (chars & tmp):
            chars |= tmp
            self.helper(arr, idx + 1, chars, res)
            chars -= tmp


if __name__ == "__main__":
    arr = [
        "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
        "o", "p"
    ]
    print(Solution().maxLength(arr))
