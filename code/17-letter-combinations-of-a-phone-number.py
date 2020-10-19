from typing import List


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        letters = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz"
        }
        res: List[str] = []
        self.helper(digits, 0, [], res, letters)
        return res

    def helper(self, digits, idx, cur, res, letters):
        if idx >= len(digits):
            tmp = [x for x in cur]
            res.append("".join(cur))
            return
        for c in letters[digits[idx]]:
            cur.append(c)
            self.helper(digits, idx + 1, cur, res, letters)
            cur.pop()


if __name__ == "__main__":
    # digits = "23"
    # digits = ""
    digits = "2"
    sol = Solution()
    print(sol.letterCombinations(digits))
