class Solution:
    def distinctEchoSubstrings(self, text: str) -> int:
        n = len(text)
        base = 29
        mod = 1000000007

        hash = [0 for _ in range(n + 1)]
        pow = [0 for _ in range(n + 1)]
        pow[0] = 1
        for i in range(1, n + 1):
            hash[i] = (hash[i - 1] * base + ord(text[i - 1])) % mod
            pow[i] = pow[i - 1] * base % mod

        def get_hash(l, r):
            nonlocal mod
            return (hash[r] - hash[l] * pow[r - l] % mod + mod) % mod

        res = set()
        for i in range(0, n):
            l = 2
            while i + l <= n:
                mid = i + l // 2
                hash1 = get_hash(i, mid)
                hash2 = get_hash(mid, i + l)
                if hash1 == hash2 and text[i:mid] == text[mid : i + l]:
                    res.add(hash1)
                l += 2
        return len(res)
