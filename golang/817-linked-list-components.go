package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	mapp := make(map[int]bool)

	res := 0
	for _, i := range nums {
		mapp[i] = true
	}

	flag := false
	for {
		if mapp[head.Val] {
			if !flag {
				flag = true
				res++
			}
		} else {
			flag = false
		}

		if head.Next == nil {
			break
		}
		head = head.Next
	}
	return res
}
