from typing import List
import heapq


class Solution:
    def kthSmallest(self, mat: List[List[int]], k: int) -> int:
        m, n = len(mat), len(mat[0])

        def get_sum(idx_list):
            return sum(row[i] for row, i in zip(mat, idx_list))

        h = []
        visited = set()
        idx = tuple(0 for _ in range(m))
        visited.add(idx)
        heapq.heappush(h, (get_sum(idx), idx))
        while k > 0:
            res, idx = heapq.heappop(h)
            k -= 1
            for i in range(m):
                if idx[i] + 1 < n:
                    new_idx = tuple(
                        idx[j] if j != i else idx[j] + 1 for j in range(m)
                    )
                    if new_idx not in visited:
                        visited.add(new_idx)
                        heapq.heappush(h, (get_sum(new_idx), new_idx))
        return res


if __name__ == "__main__":
    # mat = [[1, 3, 11], [2, 4, 6]]
    # k = 5
    # mat = [[1, 3, 11], [2, 4, 6]]
    # k = 9
    # mat = [[1, 10, 10], [1, 4, 5], [2, 3, 6]]
    # k = 7
    mat = [[1, 1, 10], [2, 2, 9]]
    k = 7
    print(Solution().kthSmallest(mat, k))
