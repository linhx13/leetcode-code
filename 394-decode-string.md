# use stack

```python
class Solution(object):
    TIMES = 1
    STRING = 2
    LEFT_PARAN = 3
    RIGHT_PARAN = 4
    
    def get_next_token(self, s, start):
        """
        :rtype: token, token_type, new-start
        """
        while start < len(s):
            if s[start] == '[':
                return '[', self.LEFT_PARAN, start+1
            elif s[start] == ']':
                return ']', self.RIGHT_PARAN, start+1
            elif s[start].isalpha():
                ret_list = []
                while start < len(s) and s[start].isalpha():
                    ret_list.append(s[start])
                    start += 1
                return ''.join(ret_list), self.STRING, start
            elif s[start].isdigit():
                ret_list = []
                while start < len(s) and s[start].isdigit():
                    ret_list.append(s[start])
                    start += 1
                return ''.join(ret_list), self.TIMES, start
        
    def decodeString(self, s):
        """
        :type s: str
        :rtype: str
        """
        stack = []
        start = 0
        while start < len(s):
            token, token_type, start = self.get_next_token(s, start)
            if token is None:
                break
            if token_type == self.TIMES:
                stack.append((token, token_type))
            elif token_type == self.STRING:
                stack.append((token, token_type))
            elif token_type == self.LEFT_PARAN:
                continue
            elif token_type == self.RIGHT_PARAN:
                string = ""
                times = 1
                while len(stack) > 0 and stack[-1][1] == self.STRING:
                    string = stack.pop()[0] + string
                if len(stack) > 0 and stack[-1][1] == self.TIMES:
                    times = int(stack.pop()[0])
                stack.append((string * times, self.STRING))
        return ''.join(map(lambda x: x[0], stack))
```

