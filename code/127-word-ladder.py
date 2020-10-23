from typing import List
from collections import deque, defaultdict


class Solution:
    def ladderLength(self, beginWord: str, endWord: str,
                     wordList: List[str]) -> int:
        word_set = set(wordList)
        if endWord not in word_set:
            return 0
        queue1, visited1 = deque(), dict()
        queue1.append(beginWord)
        visited1[beginWord] = 1
        queue2, visited2 = deque(), dict()
        queue2.append(endWord)
        visited2[endWord] = 1
        while queue1 and queue2:
            len1, len2 = len(queue1), len(queue2)
            todo = set()
            for _ in range(len1):
                s1 = queue1.popleft()
                if s1 in visited2:
                    return visited1[s1] + visited2[s1] - 1
                for w in self.gen_words(s1):
                    if w not in visited1 and w in word_set:
                        queue1.append(w)
                        visited1[w] = visited1[s1] + 1
            for _ in range(len2):
                s2 = queue2.popleft()
                if s2 in visited1:
                    return visited1[s2] + visited2[s2] - 1
                for w in self.gen_words(s2):
                    if w not in visited2 and w in word_set:
                        queue2.append(w)
                        visited2[w] = visited2[s2] + 1
        return 0

    def gen_words(self, s):
        res = []
        for i in range(len(s)):
            for c in range(ord("a"), ord("z") + 1):
                c = chr(c)
                if c != s[i]:
                    res.append(s[:i] + c + s[i+1:])
        return res


if __name__ == "__main__":
    beginWord = "hit"
    endWord = "cog"
    wordList = ["hot", "dot", "dog", "lot", "log", "cog"]
    # beginWord = "hit"
    # endWord = "cog"
    # wordList = ["hot", "dot", "dog", "lot", "log"]
    # beginWord = "a"
    # endWord = "c"
    # wordList = ["a", "b", "c"]
    print(Solution().ladderLength(beginWord, endWord, wordList))
