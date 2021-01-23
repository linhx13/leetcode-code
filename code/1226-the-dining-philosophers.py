import threading


class DiningPhilosophers:
    def __init__(self):
        self.locks = [threading.Lock() for _ in range(5)]

    # call the functions directly to execute, for example, eat()
    def wantsToEat(
        self,
        philosopher: int,
        pickLeftFork: "Callable[[], None]",
        pickRightFork: "Callable[[], None]",
        eat: "Callable[[], None]",
        putLeftFork: "Callable[[], None]",
        putRightFork: "Callable[[], None]",
    ) -> None:
        if philosopher != 0:
            first, second = philosopher, (philosopher + 1) % 5
        else:
            first, second = (philosopher + 1) % 5, philosopher
        with self.locks[first]:
            with self.locks[second]:
                pickLeftFork()
                pickRightFork()
                eat()
                putLeftFork()
                putRightFork()
