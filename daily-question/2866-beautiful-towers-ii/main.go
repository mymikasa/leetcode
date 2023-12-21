package main

func prefixSum(heights []int) ([]int64, []int64) {
	// 1, 2, 3, 4, 5, 3
	n := len(heights)
	prefix := make([]int64, n)
	suffix := make([]int64, n)
	stack1, stack2 := []int{}, []int{}

	for i := 0; i < n; i++ {
		for len(stack1) > 0 && heights[i] < heights[stack1[len(stack1)-1]] {
			stack1 = stack1[:len(stack1)-1]

		}
		if len(stack1) == 0 {
			prefix[i] = int64(i+1) * int64(heights[i])
		} else {
			last := stack1[len(stack1)-1]
			prefix[i] = prefix[last] + int64((i-last)*heights[i])
		}
		stack1 = append(stack1, i)
	}

	for i := n - 1; i >= 0; i-- {
		for len(stack2) > 0 && heights[i] < heights[stack2[len(stack2)-1]] {
			stack2 = stack2[:len(stack2)-1]
		}
		if len(stack2) == 0 {
			suffix[i] = int64(n-i) * int64(heights[i])
		} else {
			last := stack2[len(stack2)-1]
			suffix[i] = suffix[last] + int64((last-i)*heights[i])
		}
		stack2 = append(stack2, i)
	}
	return prefix, suffix
}

func maximumSumOfHeights(maxHeights []int) int64 {
	prefix, suffix := prefixSum(maxHeights)

	maxx := int64(0)
	for i := 0; i < len(maxHeights); i++ {
		maxx = max(maxx, prefix[i]+suffix[i]-int64(maxHeights[i]))
	}
	return maxx
}
