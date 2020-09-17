from threading import Lock

class H2O:
    def __init__(self):
        self.h = Lock()
        self.o = Lock()
        self.o.acquire()
        self.count = 0

    def hydrogen(self, releaseHydrogen: 'Callable[[], None]') -> None:
        self.h.acquire()
        self.count += 1
        releaseHydrogen()
        if self.count == 2:
            self.count = 0
            self.o.release()
        else:
            self.h.release()

    def oxygen(self, releaseOxygen: 'Callable[[], None]') -> None:
        self.o.acquire()
        releaseOxygen()
        self.h.release()
