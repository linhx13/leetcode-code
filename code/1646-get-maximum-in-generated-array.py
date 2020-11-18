class Solution:
    def getMaximumGenerated(self, n: int) -> int:
        if n == 0:
            return 0
        nums = [0] * (n + 1)
        nums[0] = 0
        nums[1] = 1
        for i in range(2, len(nums)):
            if i % 2 == 0:
                nums[i] = nums[i // 2]
            else:
                j = (i - 1) // 2
                nums[i] = nums[j] + nums[j + 1]
        return max(nums)


if __name__ == "__main__":
    n = 11
    print(Solution().getMaximumGenerated(n))
