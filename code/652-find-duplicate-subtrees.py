class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


from collections import Counter


class Solution:
    def findDuplicateSubtrees(self, root: TreeNode) -> List[TreeNode]:
        uids = {}
        counter = Counter()
        res = []

        def helper(node):
            if node:
                key = (node.val, helper(node.left), helper(node.right))
                if key not in uids:
                    uid = len(uids)
                    uids[key] = uid
                uid = uids[key]
                counter[uid] += 1
                if counter[uid] == 2:
                    res.append(node)
                return uid
            else:
                return None

        helper(root)
        return res
