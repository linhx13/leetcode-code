from typing import List


class Solution:
    def removeInvalidParentheses(self, s: str) -> List[str]:
        res = set()

        def dfs(idx, buf, stack, left, right):
            if idx == len(s):
                if len(stack) == 0 and left == 0 and right == 0:
                    res.add("".join(buf))
                return
            if s[idx] == "(":
                if left > 0:
                    dfs(idx + 1, buf, stack, left - 1, right)
                buf.append(s[idx])
                stack.append(s[idx])
                dfs(idx + 1, buf, stack, left, right)
                stack.pop()
                buf.pop()
            elif s[idx] == ")":
                if right > 0:
                    dfs(idx + 1, buf, stack, left, right - 1)
                if stack and stack[-1] == "(":
                    buf.append(s[idx])
                    stack.pop()
                    dfs(idx + 1, buf, stack, left, right)
                    buf.pop()
                    stack.append("(")
            else:
                buf.append(s[idx])
                dfs(idx + 1, buf, stack, left, right)
                buf.pop()

        left, right = 0, 0
        for i in range(len(s)):
            if s[i] == "(":
                left += 1
            elif s[i] == ")":
                if left > 0:
                    left -= 1
                else:
                    right += 1
        dfs(0, [], [], left, right)
        return list(res)


if __name__ == "__main__":
    s = "()())()"
    # s = "(a)())()"
    # s = ")("
    # s = "((i)"
    print(Solution().removeInvalidParentheses(s))
