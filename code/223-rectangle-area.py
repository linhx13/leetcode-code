class Solution:
    def computeArea(
        self, A: int, B: int, C: int, D: int, E: int, F: int, G: int, H: int
    ) -> int:
        w = (min(G, C) - max(E, A)) if min(G, C) > max(E, A) else 0
        h = (min(D, H) - max(B, F)) if min(D, H) > max(B, F) else 0
        return (C - A) * (D - B) + (G - E) * (H - F) - w * h
