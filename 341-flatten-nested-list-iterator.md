# python code

```python
# """
# This is the interface that allows for creating nested lists.
# You should not implement it, or speculate about its implementation
# """
#class NestedInteger(object):
#    def isInteger(self):
#        """
#        @return True if this NestedInteger holds a single integer, rather than a nested list.
#        :rtype bool
#        """
#
#    def getInteger(self):
#        """
#        @return the single integer that this NestedInteger holds, if it holds a single integer
#        Return None if this NestedInteger holds a nested list
#        :rtype int
#        """
#
#    def getList(self):
#        """
#        @return the nested list that this NestedInteger holds, if it holds a nested list
#        Return None if this NestedInteger holds a single integer
#        :rtype List[NestedInteger]
#        """

class NestedIterator(object):

    def __init__(self, nestedList):
        """
        Initialize your data structure here.
        :type nestedList: List[NestedInteger]
        """
        self.obj_stack = []
        self.idx_stack = []
        if len(nestedList) > 0:
            self.obj_stack.append(nestedList)
            self.idx_stack.append(-1)
        self._accessed_next = True  # has accessed next obj after calling hasNext()
        self._has_next = False
    
    def _next(self):
        while len(self.idx_stack) > 0:
            if self.idx_stack[-1] + 1 >= len(self.obj_stack[-1]):
                self.idx_stack.pop()
                self.obj_stack.pop()
                continue
            self.idx_stack[-1] += 1
            while self.idx_stack[-1] < len(self.obj_stack[-1]):
                obj = self.obj_stack[-1][self.idx_stack[-1]]
                if obj.isInteger():
                    self._has_next = True
                    return
                else:
                    if len(obj.getList()) == 0:
                        self.idx_stack[-1] += 1
                    else:
                        self.idx_stack.append(0)
                        self.obj_stack.append(obj.getList())
            self.idx_stack.pop()
            self.obj_stack.pop()
            continue
        self._has_next = False

    def next(self):
        """
        :rtype: int
        """
        if self._accessed_next:
            self._next()
        self._accessed_next = True
        if self._has_next:
            return self.obj_stack[-1][self.idx_stack[-1]].getInteger()
        else:
            return None

    def hasNext(self):
        """
        :rtype: bool
        """
        if self._accessed_next:
            self._next()
        self._accessed_next = False
        return self._has_next
        

# Your NestedIterator object will be instantiated and called as such:
# i, v = NestedIterator(nestedList), []
# while i.hasNext(): v.append(i.next())
```

