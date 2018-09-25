# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """

        l = ListNode(0)
        head = l
        if l1 is None:
            return l2

        if l2 is None:
            return l1


        while l1 is not None and l2 is not None:
            if l1.val < l2.val:
                l.next = l1
                l1 = l1.next
            else:
                l.next = l2
                l2 = l2.next
            l = l.next

        if l1 is None:
            l.next = l2
        else:
            l.next = l1
        return head.next


