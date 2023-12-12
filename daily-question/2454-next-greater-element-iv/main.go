package main

//func secondGreaterElement(nums []int) []int {
//	ans := make([]int, len(nums))
//	for i := range ans {
//		ans[i] = -1
//	}
//	s, t := []int{}, []int{}
//	for i, x := range nums {
//		for len(t) > 0 && nums[t[len(t)-1]] < x {
//			ans[t[len(t)-1]] = x // t 栈顶的下下个更大元素是 x
//			t = t[:len(t)-1]
//		}
//		j := len(s) - 1
//		for j >= 0 &&x < x {
//			j-- // s 栈顶的下一个更大元素是 x
//		}
//		t = append(t, s[j+1:]...) // 把从 s 弹出的这一整段元素加到 t
//		s = append(s[:j+1], i) // 当前元素（的下标）加到 s 栈顶
//	}
//	return ans
//}
//
//作者：灵茶山艾府
//链接：https://leetcode.cn/problems/next-greater-element-iv/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func secondGreaterElement(nums []int) []int {
	// 两个单调栈， 都保持单调减， 弹出一个元素代表找到 第二大元素即为此时输入的元素

	// [2,4,0,9,6]

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = -1
	}

	s, t := []int{}, []int{}

	for i, x := range nums {
		for len(t) > 0 && nums[t[len(t)-1]] < x {
			ans[t[len(t)-1]] = x
			t = t[:len(t)-1]
		}
		j := len(s) - 1

		for j >= 0 && nums[s[j]] < x {
			j--
		}

		t = append(t, s[j+1:]...)
		s = append(s[:j+1], i)
	}

	return ans
}
