from collections import Counter


class Solution:
    def minNumberOfFrogs(self, croakOfFrogs: str) -> int:
        counter = Counter()
        ans = 0
        num = 0
        for c in croakOfFrogs:
            if c == "c":
                counter["c"] += 1
                num += 1
            elif c == "r":
                if counter['c'] <= 0:
                    return -1
                counter["c"] -= 1
                counter["cr"] += 1
            elif c == "o":
                if counter['cr'] <= 0:
                    return -1
                counter["cr"] -= 1
                counter["cro"] += 1
            elif c == "a":
                if counter["cro"] <= 0:
                    return -1
                counter["cro"] -= 1
                counter["croa"] += 1
            elif c == "k":
                if counter["croa"] <= 0:
                    return -1
                counter["croa"] -= 1
                num -= 1
            ans = max(ans, num)

        if counter['c'] != 0 or counter['cr'] != 0 or counter[
                'cro'] != 0 or counter['croa'] != 0:
            return -1
        else:
            return ans


if __name__ == "__main__":
    croakOfFrogs = "croakcroak"
    print(Solution().minNumberOfFrogs(croakOfFrogs))
