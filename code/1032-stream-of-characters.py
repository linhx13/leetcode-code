from typing import List


class Node:
    def __init__(self, char):
        self.char = char
        self.final = False
        self.children = {}


class StreamChecker:
    def __init__(self, words: List[str]):
        self.root = Node("")

        def add_word(word):
            node = self.root
            for c in word:
                if c not in node.children:
                    new_node = Node(c)
                    node.children[c] = new_node
                node = node.children[c]
            node.final = True

        for word in words:
            add_word(word)
        self.nodes = set()

    def query(self, letter: str) -> bool:
        res = False
        self.nodes.add(self.root)
        new_nodes = set()
        for node in self.nodes:
            if letter in node.children:
                node = node.children[letter]
                res = node.final or res
                new_nodes.add(node)
        self.nodes = new_nodes
        return res


if __name__ == "__main__":
    words = ["ab", "ba", "aaab", "abab", "baa"]
    sc = StreamChecker(words)
    letters = "aaaaabab"
    for letter in letters:
        print(letter, sc.query(letter))
