class Solution:
    def maximumTime(self, time: str) -> str:
        def check(x, y):
            return (x[0] == "?" or x[0] == y[0]) and (
                x[1] == "?" or x[1] == y[1]
            )

        hh, mm = time.split(":")
        for x in range(23, -1, -1):
            x = ("0" + str(x)) if x < 10 else str(x)
            for y in range(59, -1, -1):
                y = ("0" + str(y)) if y < 10 else str(y)
                if check(hh, x) and check(mm, y):
                    return x + ":" + y


if __name__ == "__main__":
    # time = "2?:?0"
    # time = "0?:3?"
    # time = "1?:22"
    time = "02:00"
    print(Solution().maximumTime(time))
