from typing import List


class Node:
    def __init__(self, char):
        self.char = char
        self.children = {}
        self.idx = -1


class WordFilter:
    def __init__(self, words: List[str]):
        self.root = Node("")

        def add_str(idx, s):
            node = self.root
            for c in s:
                if c not in node.children:
                    new_node = Node(c)
                    node.children[c] = new_node
                node = node.children[c]
                node.idx = max(node.idx, idx)

        for idx, word in enumerate(words):
            for i in range(len(word)):
                s = word[i:] + "#" + word
                add_str(idx, s)

    def f(self, prefix: str, suffix: str) -> int:
        s = suffix + "#" + prefix
        node = self.root
        for c in s:
            if c not in node.children:
                return -1
            node = node.children[c]
        return node.idx


if __name__ == "__main__":
    words = ["apple", "ace"]
    wf = WordFilter(words)
    prefix, suffix = "a", "e"
    print(wf.f(prefix, suffix))
