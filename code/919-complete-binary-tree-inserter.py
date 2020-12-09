class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


from collections import deque


class CBTInserter:
    def __init__(self, root: TreeNode):
        self.root = root
        self.queue = deque()
        nodes = deque([root])
        while nodes:
            node = nodes.popleft()
            c = 0
            if node.left:
                nodes.append(node.left)
            else:
                c += 1
            if node.right:
                nodes.append(node.right)
            else:
                c += 1
            if c > 0:
                self.queue.append((node, c))

    def insert(self, v: int) -> int:
        new_node = TreeNode(v)
        node, cnt = self.queue.popleft()
        if cnt == 2:
            node.left = new_node
        else:
            node.right = new_node
        cnt -= 1
        if cnt > 0:
            self.queue.appendleft((node, cnt))
        self.queue.append((new_node, 2))
        return node.val

    def get_root(self) -> TreeNode:
        return self.root
