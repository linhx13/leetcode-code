class Solution:
    def parseBoolExpr(self, expression: str) -> bool:
        return self.parse(expression, 0, len(expression))

    def parse(self, exp, start, end):
        if exp[start] == 't':
            return True
        elif exp[start] == 'f':
            return False
        elif exp[start] == '!':
            return self.parse_not(exp, start, end)
        elif exp[start] == '&':
            return self.parse_and(exp, start, end)
        elif exp[start] == '|':
            return self.parse_or(exp, start, end)

    def parse_not(self, exp, start, end):
        return not self.parse(exp, start + 2, end - 1)

    def parse_args(self, exp, start, end):
        parens = 0
        args = []
        cur = start
        for i in range(start, end):
            if exp[i] == '(':
                parens += 1
            elif exp[i] == ')':
                parens -= 1
            elif exp[i] == ',':
                if parens == 0:
                    args.append((cur, i))
                    cur = i + 1
        if cur < end:
            args.append((cur, end))
        return args

    def parse_and(self, exp, start, end):
        args = self.parse_args(exp, start + 2, end - 1)
        vals = [self.parse(exp, x[0], x[1]) for x in args]
        return all(vals)

    def parse_or(self, exp, start, end):
        args = self.parse_args(exp, start + 2, end - 1)
        vals = [self.parse(exp, x[0], x[1]) for x in args]
        return any(vals)


if __name__ == "__main__":
    # exp = "&(t,f)"
    exp = "|(&(t,|(f,t)),!(t))"
    exp = "&(t,f,t)"
    print(Solution().parseBoolExpr(exp))
