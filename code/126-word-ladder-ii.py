from typing import List
from collections import deque


class Solution:
    def findLadders(self, beginWord: str, endWord: str,
                    wordList: List[str]) -> List[List[str]]:
        word_set = set(wordList)
        if endWord not in word_set:
            return []
        id2word = [beginWord]
        id2word.extend(list(word_set))
        word2id = {w: i for i, w in enumerate(id2word)}
        queue = deque()
        visited = set()
        queue.append([0])
        visited.add(0)
        res = []
        while queue:
            new_visited = set()
            for _ in range(len(queue)):
                cur = queue.popleft()
                word = id2word[cur[-1]]
                if word == endWord:
                    res.append([id2word[i] for i in cur])
                    continue
                for w in self.gen_words(word, word_set):
                    wid = word2id[w]
                    if wid not in visited:
                        tmp = cur[:]
                        tmp.append(wid)
                        queue.append(tmp)
                        new_visited.add(wid)
            visited.update(new_visited)
            if res:
                break
        return res

    def gen_words(self, w, word_set):
        res = []
        for i in range(len(w)):
            for c in range(ord("a"), ord("z") + 1):
                c = chr(c)
                if c != w[i]:
                    nw = w[:i] + c + w[i + 1:]
                    if nw in word_set:
                        res.append(w[:i] + c + w[i + 1:])
        return res


if __name__ == "__main__":
    beginWord = "hit"
    endWord = "cog"
    wordList = ["hot", "dot", "dog", "lot", "log", "cog"]
    # beginWord = "a"
    # endWord = "c"
    # wordList = ["a", "b", "c"]
    # beginWord = "red"
    # endWord = "tax"
    # wordList = ["ted", "tex", "red", "tax", "tad", "den", "rex", "pee"]
    print(Solution().findLadders(beginWord, endWord, wordList))
