class Solution:
    def areSentencesSimilar(self, sentence1: str, sentence2: str) -> bool:
        words1, words2 = sentence1.split(), sentence2.split()
        m, n = len(words1), len(words2)
        prefix_size = 0
        i, j = 0, 0
        while i < m and j < n:
            if words1[i] == words2[j]:
                prefix_size += 1
            else:
                break
            i, j = i + 1, j + 1
        suffix_size = 0
        i, j = m - 1, n - 1
        while i >= 0 and j >= 0:
            if words1[i] == words2[j]:
                suffix_size += 1
            else:
                break
            i, j = i - 1, j - 1
        return prefix_size + suffix_size >= min(m, n)


if __name__ == "__main__":
    sentence1 = "A A"
    sentence2 = "A aA"
    print(Solution().areSentencesSimilar(sentence1, sentence2))
