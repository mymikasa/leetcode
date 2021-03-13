# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        A = ListNode  
        tail = A(0)
        p1 = l1
        p2 = l2
        head = tail
        cnt = 0
  # 二者公共部分
        while p1 !=None and p2 !=None:
            sum = p1.val + p2.val + cnt
            cnt = int(sum/10)
            tail.next= A(sum%10)
            tail = tail.next
            p1 = p1.next
            p2 = p2.next
        
        if p1 == None:
            p1 = p2        
        while p1 !=None:

            sum =p1.val + cnt
            cnt = int(sum/10)
            tail.next = A(sum%10)
            tail = tail.next
            p1 = p1.next
        
        
        if cnt!=0:
            tail.next = A(cnt)
            tail= tail.next

        tail.next = None  
        return head.next  
