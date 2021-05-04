class Solution:
    def angleClock(self, hour: int, minutes: int) -> float:
        x = 30 * hour + 0.5 * minutes
        y = minutes * 6
        k = abs(x - y)
        return min(k, 360 - k)
