from typing import List


class Solution:
    def corpFlightBookings(self, bookings: List[List[int]],
                           n: int) -> List[int]:
        arr = [0] * (n + 3)
        for b in bookings:
            arr[b[0]] += b[2]
            arr[b[1] + 1] -= b[2]
        for i in range(1, n + 1):
            arr[i] += arr[i - 1]
        return arr[1:n + 1]
