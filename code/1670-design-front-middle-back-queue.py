from collections import deque


class FrontMiddleBackQueue:
    def __init__(self):
        self.front = deque()
        self.back = deque()

    def _adjust(self):
        total = len(self.front) + len(self.back)
        half = total - total // 2
        if len(self.front) < half:
            while len(self.front) < half:
                x = self.back.popleft()
                self.front.append(x)
        if len(self.front) > half:
            while len(self.front) > half:
                x = self.front.pop()
                self.back.appendleft(x)

    def pushFront(self, val: int) -> None:
        self.front.appendleft(val)
        self._adjust()

    def pushMiddle(self, val: int) -> None:
        total = len(self.front) + len(self.back)
        if total % 2 == 1:
            self.back.appendleft(self.front.pop())
        self.front.append(val)
        self._adjust()

    def pushBack(self, val: int) -> None:
        self.back.append(val)
        self._adjust()

    def _is_empty(self):
        total = len(self.front) + len(self.back)
        return total == 0

    def popFront(self) -> int:
        if self._is_empty():
            return -1
        if len(self.front) > 0:
            x = self.front.popleft()
        else:
            x = self.back.popleft()
        self._adjust()
        return x

    def popMiddle(self) -> int:
        if self._is_empty():
            return -1
        x = self.front.pop()
        self._adjust()
        return x

    def popBack(self) -> int:
        if self._is_empty():
            return -1
        if len(self.back) > 0:
            x = self.back.pop()
        else:
            x = self.front.pop()
        self._adjust()
        return x


if __name__ == "__main__":
    q = FrontMiddleBackQueue()
    # q.pushFront(1)
    # q.pushBack(2)
    # q.pushMiddle(3)
    # q.pushMiddle(4)
    # print(q)
    # print(q.front)
    # print(q.back)
    # print(q.popFront())
    # print(q.popMiddle())
    # print(q.popMiddle())
    # print(q.popBack())
    # print(q.popFront())
    q.pushMiddle(1)
    print(q.front, q.back)
    q.pushMiddle(2)
    print(q.front, q.back)
    q.pushMiddle(3)
    print(q.front, q.back)
    print(q.popMiddle())
    print(q.front, q.back)
    print(q.popMiddle())
    print(q.front, q.back)
    print(q.popMiddle())
