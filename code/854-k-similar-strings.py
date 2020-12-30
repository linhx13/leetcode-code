from collections import deque


class Solution:
    def kSimilarity(self, A: str, B: str) -> int:
        queue = deque()
        visited = set()
        queue.append(A)
        visited.add(A)
        n = len(A)
        res = 0
        while len(queue) > 0:
            for _ in range(len(queue)):
                s = queue.popleft()
                if s == B:
                    return res
                for i in range(n):
                    if s[i] != B[i]:
                        break
                for j in range(i + 1, n):
                    if s[j] == B[i]:
                        t = s[:i] + s[j] + s[i + 1 : j] + s[i] + s[j + 1 :]
                        if t not in visited:
                            queue.append(t)
                            visited.add(t)
            res += 1
        return res


if __name__ == "__main__":
    # A = "ab"
    # B = "ba"
    # A = "abc"
    # B = "bca"
    # A = "abac"
    # B = "baca"
    # A = "aabc"
    # B = "abca"
    A = "abccaacceecdeea"
    B = "bcaacceeccdeaae"
    # A = "abcdeabcdeabcdeabcde"
    # B = "aaaabbbbccccddddeeee"
    print(Solution().kSimilarity(A, B))
