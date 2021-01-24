from typing import List
import collections


class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        n = len(students)
        students = collections.deque(students)
        sandwiches = sandwiches[::-1]
        cnt = 0
        while len(students) > 0 and len(sandwiches) > 0:
            cur = 0
            for _ in range(len(students)):
                s = students.popleft()
                if s == sandwiches[-1]:
                    sandwiches.pop()
                    cur += 1
                else:
                    students.append(s)
            if cur == 0:
                break
            cnt += cur
        return n - cnt
