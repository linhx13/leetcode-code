class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        def get_key(s):
            return ''.join(sorted(s))

        all = set()
        l1, l2 = len(s1), len(s2)
        for i in range(l2):
            if i + l1 > l2:
                break
            all.add(get_key(s2[i:i + l1]))
        return get_key(s1) in all


if __name__ == "__main__":
    # s1 = "ab"
    # s2 = "eidbaooo"
    # s1 = "ab"
    # s2 = "eidboaoo"
    s1 = "adc"
    s2 = "dcda"
    sol = Solution()
    print(sol.checkInclusion(s1, s2))
