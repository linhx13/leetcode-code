# use stack

```python
class Solution(object):
    INTEGER = 0
    PLUS = 1
    MINUS = 2
    LEFT_PAREN = 3
    RIGHT_PAREN = 4
    
    def get_next_token(self, s, start):
        ret_list = []
        while start < len(s):
            if s[start] == " ":
                start += 1
                continue
            elif s[start] == "(":
                return "(", self.LEFT_PAREN, start+1
            elif s[start] == ")":
                return ")", self.RIGHT_PAREN, start+1
            elif s[start] == "+":
                return "+", self.PLUS, start+1
            elif s[start] == "-":
                return "-", self.MINUS, start+1
            elif s[start].isdigit():
                ret = 0
                while start < len(s) and s[start].isdigit():
                    ret = ret*10 + (ord(s[start]) - ord('0'))
                    start += 1
                return ret, self.INTEGER, start
        return None, None, start+1
            
    def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        stack = []
        start = 0
        while start < len(s):
            token, token_type, start = self.get_next_token(s, start)
            if token is None:
                break
            
            if token_type == self.INTEGER or token_type == self.PLUS or \
                token_type == self.MINUS or token_type == self.LEFT_PAREN:
                stack.append(token)
            elif token_type == self.RIGHT_PAREN:
                r = 0
                to_find_oper = False
                b = 0
                while len(stack) > 0 and stack[-1] != '(':
                    if to_find_oper:
                        oper = stack.pop()
                        to_find_oper = False
                        if oper == "+":
                            r += b
                        elif oper == "-":
                            r -= b
                    else:
                        b = stack.pop()
                        to_find_oper = True
                if len(stack) > 0 and stack[-1] == '(':
                    stack.pop()
                r += b
                stack.append(r)
        r = 0
        b = 0
        to_find_oper = False
        while len(stack) > 0:
            if to_find_oper:
                oper = stack.pop()
                to_find_oper = False
                if oper == "+":
                    r += b
                elif oper == "-":
                    r -= b
            else:
                b = stack.pop()
                to_find_oper = True
        r += b
        return r
                    
```

