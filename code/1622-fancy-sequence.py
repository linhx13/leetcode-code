class Fancy:
    def __init__(self):
        self.data = []
        self.add = [0]
        self.mul = [1]
        self.mod = int(10**9 + 7)

    def append(self, val: int) -> None:
        self.data.append(val)
        self.add.append(self.add[-1])
        self.mul.append(self.mul[-1])

    def addAll(self, inc: int) -> None:
        self.add[-1] += inc

    def multAll(self, m: int) -> None:
        self.add[-1] = self.add[-1] * m % self.mod
        self.mul[-1] = self.mul[-1] * m % self.mod

    def getIndex(self, idx: int) -> int:
        if idx >= len(self.data):
            return -1
        m = self.mul[-1] * pow(self.mul[idx], self.mod - 2, self.mod)
        inc = self.add[-1] - self.add[idx] * m
        return (self.data[idx] * m + inc) % self.mod
