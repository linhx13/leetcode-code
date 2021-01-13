class Solution:
    def halvesAreAlike(self, s: str) -> bool:
        vowels = {"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
        a = sum(1 for c in s[: len(s) // 2] if c in vowels)
        b = sum(1 for c in s[len(s) // 2 :] if c in vowels)
        return a == b
