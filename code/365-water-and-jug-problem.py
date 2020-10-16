class Solution:
    def canMeasureWater(self, x: int, y: int, z: int) -> bool:
        if x + y < z:
            return False
        if z == y or z == x or z == 0 or z == x + y:
            return True
        x, y = min(x, y), max(x, y)
        return not (z % self.gcd(x, y))

    def gcd(self, x, y):
        if x:
            return self.gcd(y % x, x)
        else:
            return y
