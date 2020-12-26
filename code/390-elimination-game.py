class Solution:
    def lastRemaining(self, n: int) -> int:
        return self.left_to_right(n)

    def left_to_right(self, n):
        if n <= 2:
            return n
        return 2 * self.right_to_left(n // 2)

    def right_to_left(self, n):
        if n <= 2:
            return 1
        if n % 2 == 1:
            return 2 * self.left_to_right(n // 2)
        return 2 * self.left_to_right(n // 2) - 1
