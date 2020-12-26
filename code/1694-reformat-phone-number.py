class Solution:
    def reformatNumber(self, number: str) -> str:
        number = number.replace("-", "").replace(" ", "")
        remaining = len(number)
        res = []
        idx = 0
        while remaining > 0:
            if remaining > 4:
                res.append(number[idx : idx + 3])
                idx += 3
                remaining -= 3
            elif remaining == 4:
                res.append(number[idx : idx + 2])
                res.append(number[idx + 2 : idx + 4])
                remaining -= 4
                idx += 4
            elif remaining == 3:
                res.append(number[idx : idx + 3])
                remaining -= 3
                idx += 3
            elif remaining == 2:
                res.append(number[idx : idx + 2])
                remaining -= 2
                idx += 2
        return "-".join(res)
