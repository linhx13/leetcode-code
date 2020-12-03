from typing import Optional, List, Union
import random


class Node:
    def __init__(self, data: Optional[int] = None):
        self.data = data
        self.forward: Optional[List] = None


class Skiplist:
    _MAX_LEVEL = 16

    def __init__(self):
        self.level = 1
        self.head = Node()
        self.head.forward: List = [None] * Skiplist._MAX_LEVEL

    def search(self, target: int) -> bool:
        p = self.head
        for i in range(self.level - 1, -1, -1):
            while p.forward[i] and p.forward[i].data < target:
                p = p.forward[i]
        return p.forward[0] and p.forward[0].data == target

    def _random_level(self, p: float = 0.5) -> int:
        level = 1
        while level < Skiplist._MAX_LEVEL and random.random() < p:
            level += 1
        return level

    def add(self, num: int) -> None:
        level = self._random_level()
        if self.level < level:
            self.level = level
        node = Node(num)
        node.forward = [None] * level
        update = [self.head] * level
        p = self.head
        for i in range(level - 1, -1, -1):
            while p.forward[i] and p.forward[i].data < num:
                p = p.forward[i]
            update[i] = p
        for i in range(level):
            node.forward[i] = update[i].forward[i]
            update[i].forward[i] = node

    def erase(self, num: int) -> bool:
        update: List = [None] * self.level
        p = self.head
        for i in range(self.level - 1, -1, -1):
            while p.forward[i] and p.forward[i].data < num:
                p = p.forward[i]
            update[i] = p
        if p.forward[0] and p.forward[0].data == num:
            for i in range(self.level - 1, -1, -1):
                if update[i].forward[i] == p.forward[0]:
                    update[i].forward[i] = update[i].forward[i].forward[i]
            return True
        else:
            return False
