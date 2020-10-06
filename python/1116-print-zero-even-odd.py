from threading import Lock

class ZeroEvenOdd:
    def __init__(self, n):
        self.n = n
        self.lock_zero = Lock()
        self.lock_even = Lock()
        self.lock_odd = Lock()
        self.lock_odd.acquire()
        self.lock_even.acquire()

    # printNumber(x) outputs "x", where x is an integer.
    def zero(self, printNumber: 'Callable[[int], None]') -> None:
        for i in range(1, self.n+1):
            self.lock_zero.acquire()
            printNumber(0)
            if i % 2 == 0:
                self.lock_even.release()
            else:
                self.lock_odd.release()

    def even(self, printNumber: 'Callable[[int], None]') -> None:
        for i in range(2, self.n+1, 2):
            self.lock_even.acquire()
            printNumber(i)
            self.lock_zero.release()

    def odd(self, printNumber: 'Callable[[int], None]') -> None:
        for i in range(1, self.n+1, 2):
            self.lock_odd.acquire()
            printNumber(i)
            self.lock_zero.release()


def printNumber(x):
    import sys
    print(x, file=sys.stdout)

def main():
    from functools import partial
    import threading
    obj = ZeroEvenOdd(4)
    t0 = threading.Thread(target=partial(obj.zero, printNumber))
    t0.start()
    t1 = threading.Thread(target=partial(obj.odd, printNumber))
    t1.start()
    t2 = threading.Thread(target=partial(obj.even, printNumber))
    t2.start()
    t0.join()
    t1.join()
    t2.join()


if __name__ == '__main__':
    main()
