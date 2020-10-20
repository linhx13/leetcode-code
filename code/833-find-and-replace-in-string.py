from typing import List


class Solution:
    def findReplaceString(self, S: str, indexes: List[int], sources: List[str],
                          targets: List[str]) -> str:
        if len(indexes) == 0:
            return S
        items = []
        for i in range(len(indexes)):
            items.append((indexes[i], sources[i], targets[i]))
        items.sort(key=lambda x: x[0])
        arr = []
        arr.append(S[:items[0][0]])
        for i in range(len(items)):
            if i == len(items) - 1:
                part = S[items[i][0]:]
            else:
                part = S[items[i][0]:items[i + 1][0]]
            if part.startswith(items[i][1]):
                part = items[i][2] + part[len(items[i][1]):]
            arr.append(part)
        return "".join(arr)


if __name__ == "__main__":
    S = "vmokgggqzp"
    indexes = [3, 5, 1]
    sources = ["kg", "ggq", "mo"]
    targets = ["s", "so", "bfr"]
    sol = Solution()
    print(sol.findReplaceString(S, indexes, sources, targets))
