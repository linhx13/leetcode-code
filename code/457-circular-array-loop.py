from typing import List


class Solution:
    def circularArrayLoop(self, nums: List[int]) -> bool:
        def get_next(i):
            n = len(nums)
            return (i + nums[i]) % n

        for i in range(len(nums)):
            if nums[i] == 0:
                continue
            j = i
            k = get_next(i)
            while nums[k] * nums[i] > 0 and nums[get_next(k)] * nums[i] > 0:
                if j == k:
                    if j == get_next(j):
                        break
                    return True
                j = get_next(j)
                k = get_next(get_next(k))
            j = i
            val = nums[i]
            while nums[j] * val > 0:
                k = get_next(j)
                nums[j] = 0
                j = k
        return False


if __name__ == "__main__":
    nums = [-1, 2]
    print(Solution().circularArrayLoop(nums))
