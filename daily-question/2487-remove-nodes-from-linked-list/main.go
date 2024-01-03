package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	res := &ListNode{}
	for head != nil {
		p := head
		head = head.Next
		p.Next = res.Next
		res.Next = p
	}
	return res.Next
}

func removeNodesV1(head *ListNode) *ListNode {
	// 翻转链表

	//s := []int{}
	//s = append(s, head.Val)
	//minn := head.Val
	//for head.Next != nil {
	//	head = head.Next
	//	v := head.Val
	//	minn = min(v, minn)
	//
	//	if v > minn {
	//		for s[len(s)-1] < v {
	//			s = s[:len(s)-1]
	//		}
	//	}
	//}

	head = reverse(head)

	// [5,2,13,3,8]
	for p := head; p.Next != nil; {
		if p.Val > p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return reverse(head)
}

// 5 3 2 13 3 8
func removeNodesV2(head *ListNode) *ListNode {
	// 单调栈
	s := []*ListNode{}
	s = append(s, head)
	//minn := head.Val
	for head.Next != nil {
		head = head.Next
		for len(s) > 0 && head.Val > s[len(s)-1].Val {
			s = s[:len(s)-1]
		}
		s = append(s, head)
	}
	for i := 0; i < len(s)-1; i++ {
		s[i].Next = s[i+1]
	}

	return s[0]
}

func removeNodes(head *ListNode) *ListNode {
	// 递归
	if head == nil {
		return nil
	}

	head.Next = removeNodes(head.Next)

	if head.Next != nil && head.Val < head.Next.Val {
		return head.Next
	} else {
		return head
	}
}
