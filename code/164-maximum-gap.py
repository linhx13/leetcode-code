from typing import List


class Bucket:
    def __init__(self):
        self.used = False
        self.min_val = float("inf")
        self.max_val = float("-inf")


class Solution:
    def maximumGap(self, nums: List[int]) -> int:
        if (not nums) or len(nums) < 2:
            return 0

        mini = min(nums)
        maxi = max(nums)

        bucket_size = max(1, (maxi - mini) // (len(nums) - 1))
        bucket_num = (maxi - mini) // bucket_size + 1
        buckets = [Bucket() for _ in range(bucket_num)]

        for n in nums:
            idx = (n - mini) // bucket_size
            buckets[idx].used = True
            buckets[idx].min_val = min(n, buckets[idx].min_val)
            buckets[idx].max_val = max(n, buckets[idx].max_val)

        prev_max = mini
        max_gap = 0
        for bucket in buckets:
            if not bucket.used:
                continue
            max_gap = max(max_gap, bucket.min_val - prev_max)
            prev_max = bucket.max_val
        return max_gap
