class Solution:
    def checkOverlap(
        self,
        radius: int,
        x_center: int,
        y_center: int,
        x1: int,
        y1: int,
        x2: int,
        y2: int,
    ) -> bool:
        def helper(val, mi, ma):
            return max(mi, min(val, ma))

        cx = helper(x_center, x1, x2)
        cy = helper(y_center, y1, y2)

        dist = (cx - x_center) ** 2 + (cy - y_center) ** 2
        return dist <= (radius ** 2)
