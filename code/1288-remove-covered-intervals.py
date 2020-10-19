from typing import List


class Solution:
    def removeCoveredIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort(key=lambda x: (x[0], -x[1]))
        res = 0
        i = 0
        while i < len(intervals):
            res += 1
            j = i + 1
            while j < len(intervals):
                if not (intervals[i][0] <= intervals[j][0]
                        and intervals[j][1] <= intervals[i][1]):
                    break
                j += 1
            i = j
        return res


if __name__ == "__main__":
    intervals = [[1, 4], [3, 6], [2, 8]]
    # intervals = [[1, 4], [2, 3]]
    # intervals = [[0, 10], [5, 12]]
    # intervals = [[3, 10], [4, 10], [5, 11]]
    # intervals = [[1, 2], [1, 4], [3, 4]]
    sol = Solution()
    print(sol.removeCoveredIntervals(intervals))
