from typing import List


class Solution:
    def solveEquation(self, equation: str) -> str:
        arr = equation.split("=")
        x1, v1 = self.parse(arr[0])
        x2, v2 = self.parse(arr[1])
        if x1 == x2:
            if v1 == v2:
                return "Infinite solutions"
            else:
                return "No solution"
        return "x={}".format((v2 - v1) // (x1 - x2))

    def parse(self, A):
        x, v = 0, 0
        i = 0
        while i < len(A):
            sign = 1
            num = 0
            if A[i] == "-":
                sign = -1
                i += 1
            elif A[i] == "+":
                i += 1
            if A[i] == "x":
                x += sign
                i += 1
                continue
            while i < len(A) and A[i].isdigit():
                num = num * 10 + ord(A[i]) - ord("0")
                i += 1
            if i == len(A):
                v += sign * num
            elif A[i] == "x":
                x += sign * num
                i += 1
            else:
                v += sign * num
        return x, v


if __name__ == "__main__":
    # equation = "x+5-3+x=6+x-2"
    # equation = "2x+3x-6x=x+2"
    equation = "x=x+2"
    sol = Solution()
    print(sol.solveEquation(equation))
