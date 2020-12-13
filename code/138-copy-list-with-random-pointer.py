class Node:
    def __init__(self, x: int, next: "Node" = None, random: "Node" = None):
        self.val = int(x)
        self.next = next
        self.random = random


class Solution:
    def copyRandomList(self, head: "Node") -> "Node":
        mapping = {}

        def get_node(node):
            if node in mapping:
                return mapping[node]
            res = Node(node.val)
            mapping[node] = res
            return res

        old = head
        cur = dummy = Node(0)
        while old is not None:
            new = get_node(old)
            if old.random is not None:
                new.random = get_node(old.random)
            cur.next = new
            cur = new
            old = old.next
        return dummy.next
