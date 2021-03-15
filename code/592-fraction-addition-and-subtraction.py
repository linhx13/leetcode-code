import math


class Solution:
    def fractionAddition(self, expression: str) -> str:
        i, j, a, b, n = 0, 0, 0, 1, len(expression)
        while j < n:
            i, j = j, expression.find("/", j)
            a2 = int(expression[i:j])
            i, j = j + 1, min(
                expression.find("+", j) % (n + 1),
                expression.find("-", j) % (n + 1),
            )
            b2 = int(expression[i:j])
            a, b = a * b2 + b * a2, b * b2
            gcd = math.gcd(a, b)
            a, b = a // gcd, b // gcd
        return str(a) + "/" + str(b)
