# python code

``` python
class Solution(object):
    def isValidSerialization(self, preorder):
        """
        :type preorder: str
        :rtype: bool
        """
        stack = [] 
        input_list = preorder.split(',')
        for item in input_list:
            stack.append(item)
            while len(stack) > 2 and stack[-1] == '#' and stack[-2] == '#' and stack[-3] != '#':
                stack.pop()
                stack.pop()
                stack[-1] = '#'
        if len(stack) == 1 and stack[0] == '#':
            return True
        else:
            return False
```

