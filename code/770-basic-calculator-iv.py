from typing import List
import collections


class Poly(collections.Counter):
    def __add__(self, other):
        self.update(other)
        return self

    def __sub__(self, other):
        self.update({k: -v for k, v in other.items()})
        return self

    def __mul__(self, other):
        res = Poly()
        for k1, v1 in self.items():
            for k2, v2 in other.items():
                res.update({tuple(sorted(k1 + k2)): v1 * v2})
        return res

    def eval(self, evalmap):
        res = Poly()
        for k, v in self.items():
            free = []
            for token in k:
                if token in evalmap:
                    v *= evalmap[token]
                else:
                    free.append(token)
            res[tuple(free)] += v
        return res

    def to_list(self):
        return [
            "*".join((str(v),) + k)
            for k, v in sorted(
                self.items(), key=lambda x: (-len(x[0]), x[0], x[1])
            )
            if v
        ]


class Solution:
    def basicCalculatorIV(
        self, expression: str, evalvars: List[str], evalints: List[int]
    ) -> List[str]:
        def combine(symbol, left, right):
            if symbol == "+":
                return left + right
            if symbol == "-":
                return left - right
            if symbol == "*":
                return left * right

        def make(expr: str):
            res = Poly()
            if expr.isdigit():
                res.update({(): int(expr)})
            else:
                res[(expr,)] += 1
            return res

        def parse(expr):
            bucket = []
            symbols = []
            i = 0
            while i < len(expr):
                if expr[i] == "(":
                    bal = 0
                    for j in range(i, len(expr)):
                        if expr[j] == "(":
                            bal += 1
                        if expr[j] == ")":
                            bal -= 1
                        if bal == 0:
                            break
                    bucket.append(parse(expr[i + 1 : j]))
                    i = j
                elif expr[i].isalnum():
                    for j in range(i, len(expr)):
                        if expr[j] == " ":
                            bucket.append(make(expr[i:j]))
                            break
                    else:
                        bucket.append(make(expr[i:]))
                    i = j
                elif expr[i] in "+-*":
                    symbols.append(expr[i])
                i += 1

            for i in range(len(symbols) - 1, -1, -1):
                if symbols[i] == "*":
                    bucket[i] = combine(
                        symbols.pop(i), bucket[i], bucket.pop(i + 1)
                    )

            if not bucket:
                return Poly()
            res = bucket[0]
            for i, symbol in enumerate(symbols):
                res = combine(symbol, res, bucket[i + 1])
            return res

        evalmap = dict(zip(evalvars, evalints))
        res = parse(expression).eval(evalmap)
        return res.to_list()


if __name__ == "__main__":
    # expression = "a * b * c + b * a * c * 4"
    # evalvars = []
    # evalints = []
    expression = "e - 8 + temperature - pressure"
    evalvars = ["e", "temperature"]
    evalints = [1, 12]
    print(Solution().basicCalculatorIV(expression, evalvars, evalints))
