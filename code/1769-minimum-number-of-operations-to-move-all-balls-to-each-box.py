from typing import List


class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        left, right = 0, 0
        res = [0] * len(boxes)
        for i, x in enumerate(boxes):
            if x == "1":
                right += 1
                res[0] += i
        if boxes[0] == "1":
            right -= 1
            left += 1
        for i in range(1, len(boxes)):
            res[i] = res[i - 1] + left - right
            if boxes[i] == "1":
                right -= 1
                left += 1
        return res


if __name__ == "__main__":
    boxes = "001011"
    print(Solution().minOperations(boxes))
