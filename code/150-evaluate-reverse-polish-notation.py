from typing import List


class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        ops = set(["+", "-", "*", "/"])
        for t in tokens:
            if t not in ops:
                stack.append(int(t))
            else:
                y = stack.pop()
                x = stack.pop()
                if t == "+":
                    stack.append(x + y)
                elif t == "-":
                    stack.append(x - y)
                elif t == "*":
                    stack.append(x * y)
                elif t == "/":
                    stack.append(int(x / y))
        return stack[0]
