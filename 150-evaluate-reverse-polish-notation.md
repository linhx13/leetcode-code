# use stack

```python
class Solution(object):
    def evalRPN(self, tokens):
        """
        :type tokens: List[str]
        :rtype: int
        """
        stack = []
        for token in tokens:
            if token in ["+", "-", "*", "/"]:
                b = stack.pop()
                a = stack.pop()
                if token == '+':
                    c = a + b
                elif token == '-':
                    c = a - b
                elif token == '*':
                    c = a * b
                else:
                    c = int(1.0 * a / b)
                stack.append(c)
            else:
                stack.append(int(token))
        return stack[0]
```

