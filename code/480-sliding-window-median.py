from typing import List


class Solution:
    def medianSlidingWindow(self, nums: List[int], k: int) -> List[float]:
        if not nums:
            return []
        window = [i for i in range(k)]
        window.sort(key=lambda x: nums[x])
        d = {x: i for i, x in enumerate(window)}

        def get_median():
            i = k // 2
            if k % 2 == 0:
                return (nums[window[i]] + nums[window[i - 1]]) / 2
            else:
                return nums[window[i]]

        res = []
        res.append(get_median())

        for i in range(k, len(nums)):
            j = d[i - k]
            del d[i - k]
            window[j] = i
            d[i] = j
            while j + 1 < k and nums[window[j]] > nums[window[j + 1]]:
                d[window[j]] = j + 1
                d[window[j + 1]] = j
                window[j], window[j + 1] = window[j + 1], window[j]
                j += 1
            while j - 1 >= 0 and nums[window[j]] < nums[window[j - 1]]:
                d[window[j]] = j - 1
                d[window[j - 1]] = j
                window[j], window[j - 1] = window[j - 1], window[j]
                j -= 1
            res.append(get_median())
        return res


if __name__ == "__main__":
    # nums = [1, 3, -1, -3, 5, 3, 6, 7]
    # k = 3
    # nums = [1, 4, 2, 3]
    # k = 4
    nums = [7, 9, 3, 8, 0, 2, 4, 8, 3, 9]
    k = 1
    print(Solution().medianSlidingWindow(nums, k))
