from typing import List


class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        m, n = len(board), len(board[0])
        prefix_dict = {}
        for w in words:
            for i in range(1, len(w)):
                key = w[:i]
                if key not in prefix_dict:
                    prefix_dict[key] = '$'
            prefix_dict[w] = w

        res = set()
        seen = set()

        def dfs(r, c, cur):
            if not (0 <= r < m and 0 <= c < n and (r, c) not in seen):
                return
            cur = cur + board[r][c]
            val = prefix_dict.get(cur)
            if val is None:
                return
            seen.add((r, c))
            if val != '$':
                res.add(val)
            dfs(r - 1, c, cur)
            dfs(r + 1, c, cur)
            dfs(r, c - 1, cur)
            dfs(r, c + 1, cur)
            seen.remove((r, c))

        for r in range(m):
            for c in range(n):
                dfs(r, c, "")

        return list(res)


if __name__ == "__main__":
    # board = [["o", "a", "a", "n"], ["e", "t", "a", "e"], ["i", "h", "k", "r"],
    #          ["i", "f", "l", "v"]]
    # words = ["oath", "pea", "eat", "rain", "oat", "r"]
    board = [["a", "a"]]
    words = ["a"]
    print(Solution().findWords(board, words))
