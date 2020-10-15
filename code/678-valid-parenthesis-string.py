class Solution:
    def checkValidString(self, s: str) -> bool:
        left, right = 0, 0
        for c in s:
            if c == "(":
                left += 1
                right += 1
            elif c == ")":
                left = max(left - 1, 0)
                right -= 1
            elif c == "*":
                left = max(left - 1, 0)
                right += 1
            if right < 0:
                return False
        return left <= 0 and 0 <= right


if __name__ == "__main__":
    s = "(*))"
    sol = Solution()
    print(sol.checkValidString(s))
