class Solution:
    def evaluate(self, expression: str) -> int:
        def get_token():
            nonlocal idx, scopes
            while idx < len(expression) and expression[idx] == " ":
                idx += 1
            if idx == len(expression):
                return None
            if expression[idx] == "(":
                idx += 1
                return "("
            if expression[idx] == ")":
                idx += 1
                return ")"
            start = idx
            while idx < len(expression) and expression[idx] not in [
                " ",
                "(",
                ")",
            ]:
                idx += 1
            return expression[start:idx]

        def get_val(token):
            nonlocal scopes
            try:
                return int(token)
            except:
                for i in range(len(scopes) - 1, -1, -1):
                    if token in scopes[i]:
                        return scopes[i][token]
                return None

        def parse_expr() -> int:
            token = get_token()
            if token == "(":
                return mapping[get_token()]()
            else:
                return get_val(token)

        def wrap(fn):
            nonlocal scopes

            def _fn(*args, **kwargs):
                scopes.append({})
                res = fn(*args, **kwargs)
                scopes.pop()
                return res

            return _fn

        @wrap
        def parse_add() -> int:
            x = parse_expr()
            y = parse_expr()
            assert ")" == get_token()
            return x + y

        @wrap
        def parse_mult() -> int:
            x = parse_expr()
            y = parse_expr()
            assert ")" == get_token()
            return x * y

        @wrap
        def parse_let() -> int:
            var, val = None, None
            while True:
                token = get_token()
                if token == ")":
                    break
                if token == "(":
                    val = mapping[get_token()]()
                else:
                    if var is None:
                        var = token
                    else:
                        val = get_val(token)
                if var is not None and val is not None:
                    scopes[-1][var] = val
                    var, val = None, None
            if val is None:
                val = get_val(var)
            return val

        mapping = {"add": parse_add, "mult": parse_mult, "let": parse_let}
        scopes = []
        idx = 0

        return parse_expr()


if __name__ == "__main__":
    # expression = "(add 1 2)"
    # expression = "(mult 3 (add 2 3))"
    # expression = "(let x 2 (mult x 5))"
    # expression = "(let x 2 (mult x (let x 3 y 4 (add x y))))"
    # expression = "(let x 3 x 2 x)"
    # expression = "(let x 1 y 2 x (add x y) (add x y))"
    # expression = "(let x 2 (add (let x 3 (let x 4 x)) x))"
    expression = "(let a1 3 b2 (add a1 1) b2) "
    print(Solution().evaluate(expression))
