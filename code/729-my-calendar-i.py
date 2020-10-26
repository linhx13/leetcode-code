class MyCalendar:
    def __init__(self):
        self.events = []

    def book(self, start: int, end: int) -> bool:
        for i in range(len(self.events)):
            if (self.events[i][0] <= start < self.events[i][1]
                or self.events[i][0] < end <= self.events[i][1]) \
               or (start <= self.events[i][0] < end
                  or start < self.events[i][1] <= end):
                return False
        self.events.append((start, end))
        return True
