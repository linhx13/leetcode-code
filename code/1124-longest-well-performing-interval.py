from typing import List


class Solution:
    def longestWPI(self, hours: List[int]) -> int:
        sums = [0]
        if hours[0] > 8:
            sums.append(1)
        else:
            sums.append(-1)
        for i in range(1, len(hours)):
            if hours[i] > 8:
                sums.append(sums[-1] + 1)
            else:
                sums.append(sums[-1] - 1)
        stack = [0]
        res = 0
        for i in range(1, len(sums)):
            if sums[stack[-1]] > sums[i]:
                stack.append(i)
        for i in range(len(sums) - 1, -1, -1):
            if stack and stack[-1] == i:
                stack.pop()
            while stack and sums[stack[-1]] < sums[i]:
                res = max(res, i - stack[-1])
                stack.pop()
        return res


if __name__ == "__main__":
    hours = [9, 9, 6, 6, 6, 9, 9, 6, 6, 9]
    sol = Solution()
    print(sol.longestWPI(hours))
