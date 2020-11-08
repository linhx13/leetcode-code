from typing import List


class Solution:
    def rankTeams(self, votes: List[str]) -> str:
        rank = [[0 for _ in range(len(votes[0]))] for _ in range(26)]
        for v in votes:
            for i, c in enumerate(v):
                rank[ord(c) - ord('A')][i] -= 1
        res = votes[0]
        res = sorted(res, key=lambda x: (rank[ord(x) - ord('A')], x))
        return "".join(res)


if __name__ == "__main__":
    votes = ["ABC", "ACB", "ABC", "ACB", "ACB"]
    print(Solution().rankTeams(votes))
