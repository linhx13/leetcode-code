from typing import Set, Optional


class Block:
    def __init__(self, val=0):
        self.val: int = val
        self.keys: Set = set()
        self.prev: Optional[Block] = None
        self.next: Optional[Block] = None

    def remove(self):
        self.prev.next = self.next
        self.next.prev = self.prev
        self.prev, self.next = None, None

    def insert_after(self, new_block):
        old_next = self.next
        self.next = new_block
        new_block.prev = self
        new_block.next = old_next
        old_next.prev = new_block


class AllOne:
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.head = Block()
        self.tail = Block()
        self.head.next = self.tail
        self.tail.prev = self.head
        self.mapping = {}

    def inc(self, key: str) -> None:
        """
        Inserts a new key <Key> with value 1. Or increments an existing key by 1.
        """
        if key not in self.mapping:
            cur_block = self.head
        else:
            cur_block = self.mapping[key]
            cur_block.keys.remove(key)

        if cur_block.val + 1 != cur_block.next.val:
            new_block = Block(cur_block.val + 1)
            cur_block.insert_after(new_block)
        else:
            new_block = cur_block.next
        new_block.keys.add(key)
        self.mapping[key] = new_block

        if not cur_block.keys and cur_block.val != 0:
            cur_block.remove()

    def dec(self, key: str) -> None:
        """
        Decrements an existing key by 1. If Key's value is 1, remove it from the data structure.
        """
        if key not in self.mapping:
            return
        cur_block = self.mapping[key]
        del self.mapping[key]
        cur_block.keys.remove(key)

        if cur_block.val != 1:
            if cur_block.val - 1 != cur_block.prev.val:
                new_block = Block(cur_block.val - 1)
                cur_block.prev.insert_after(new_block)
            else:
                new_block = cur_block.prev
            new_block.keys.add(key)
            self.mapping[key] = new_block

        if not cur_block.keys:
            cur_block.remove()

    def getMaxKey(self) -> str:
        """
        Returns one of the keys with maximal value.
        """
        if self.tail.prev.val == 0:
            return ""
        return next(iter(self.tail.prev.keys))

    def getMinKey(self) -> str:
        """
        Returns one of the keys with Minimal value.
        """
        if self.head.next.val == 0:
            return ""
        return next(iter(self.head.next.keys))
