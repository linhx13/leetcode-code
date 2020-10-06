from threading import Semaphore

class FizzBuzz:
    def __init__(self, n: int):
        self.n = n
        self.sn = Semaphore(0)
        self.s3 = Semaphore(0)
        self.s5 = Semaphore(0)
        self.s15 = Semaphore(0)
        # self.s3.acquire()
        # self.s5.acquire()
        # self.s15.acquire()

    # printFizz() outputs "fizz"
    def fizz(self, printFizz: 'Callable[[], None]') -> None:
        for i in range(3, self.n+1, 3):
            if i % 5 == 0:
                continue
            self.s3.acquire()
            printFizz()
            self.sn.release()

    # printBuzz() outputs "buzz"
    def buzz(self, printBuzz: 'Callable[[], None]') -> None:
        for i in range(5, self.n+1, 5):
            if i % 3 == 0:
                continue
            self.s5.acquire()
            printBuzz()
            self.sn.release()

    # printFizzBuzz() outputs "fizzbuzz"
    def fizzbuzz(self, printFizzBuzz: 'Callable[[], None]') -> None:
        for i in range(15, self.n+1, 15):
            self.s15.acquire()
            printFizzBuzz()
            self.sn.release()

    # printNumber(x) outputs "x", where x is an integer.
    def number(self, printNumber: 'Callable[[int], None]') -> None:
        for i in range(1, self.n+1):
            if i % 15 == 0:
                self.s15.release()
                self.sn.acquire()
            elif i % 5 == 0:
                self.s5.release()
                self.sn.acquire()
            elif i % 3 == 0:
                self.s3.release()
                self.sn.acquire()
            else:
                printNumber(i)
