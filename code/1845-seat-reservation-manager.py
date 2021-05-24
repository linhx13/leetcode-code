import heapq


class SeatManager:
    def __init__(self, n: int):
        self.arr = list(range(1, n + 1))
        heapq.heapify(self.arr)

    def reserve(self) -> int:
        return heapq.heappop(self.arr)

    def unreserve(self, seatNumber: int) -> None:
        heapq.heappush(self.arr, seatNumber)
