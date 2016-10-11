# two pointers

```python
class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        i = 0
        j = len(height) - 1
        ret = 0
        while i < j:
            if height[i] < height[j]:
                h = height[i]
            else:
                h = height[j]
            area = (j - i) * h
            if area > ret:
                ret = area
            if height[i] < height[j]:
                i += 1
            else:
                j -= 1
        return ret
        
```

