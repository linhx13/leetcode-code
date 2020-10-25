class Solution:
    def lengthLongestPath(self, input: str) -> int:
        stack = []
        res = 0
        for path in input.split("\n"):
            subs = len(path) - len(path.lstrip('\t'))
            while len(stack) > subs:
                stack.pop()
            stack.append(path.lstrip("\t"))
            if "." in path:
                filename = "/".join(stack)
                res = max(res, len(filename))
        return res
