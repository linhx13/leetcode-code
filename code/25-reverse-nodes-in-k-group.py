class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        c = 0
        res = ptr = ListNode(0)
        cur = head
        while head:
            c += 1
            if c == k:
                next, head.next = head.next, None
                ptr.next = self.reverse_list(cur)
                ptr = cur
                c = 0
                head = cur = next
            else:
                head = head.next
        if cur:
            ptr.next = cur
        return res.next

    def reverse_list(self, head):
        res = ptr = ListNode(0)
        while head:
            next = head.next
            ptr.next, head.next = head, ptr.next
            head = next
        return ptr.next


def arr2list(arr):
    res = ptr = ListNode(0)
    for i in arr:
        ptr.next = ListNode(i)
        ptr = ptr.next
    return res.next


def list2arr(head):
    res = []
    while head:
        res.append(head.val)
        head = head.next
    return res


if __name__ == "__main__":
    # head = [1, 2, 3, 4, 5]
    head = [1]
    k = 3
    data = arr2list(head)
    res = Solution().reverseKGroup(data, k)
    print(list2arr(res))
