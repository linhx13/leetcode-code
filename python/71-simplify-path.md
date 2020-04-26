# use stack directly

```python
class Solution(object):
    def simplifyPath(self, path):
        """
        :type path: str
        :rtype: str
        """
        stack = []
        for item in path.strip().split('/'):
            if item == ".":
                continue
            elif item == "..":
                if len(stack) > 0:
                    stack.pop()
            elif item == "":
                continue
            else:
                stack.append(item)
        return '/%s' % '/'.join(stack)
```

