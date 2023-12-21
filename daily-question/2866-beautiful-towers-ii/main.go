package main

func prefixSum(heights []int) ([]int64, []int64) {
	// 1, 2, 3, 4, 5, 3
	n := len(heights)
	prefix := make([]int64, n)
	suffix := make([]int64, n)
	stack1, _ := []int{}, []int{}
	for i := 0; i < len(heights); i++ {
		for len(stack1) > 0 && heights[i] < heights[stack1[len(stack1)-1]] {
			stack1 = stack1[:len(stack1)-1]
		}

		if len(stack1) == 0 {
			prefix[i] = int64(i+1) * int64(heights[i])
		} else {
			last := stack1[len(stack1)-1]
			prefix[i] = prefix[last] + int64((len(stack1)-i+1)*heights[i])
		}
		stack1 = append(stack1, i)
	}
	return prefix, suffix
}

func maximumSumOfHeights(maxHeights []int) int64 {
	prefix, suffix := prefixSum(maxHeights)

	maxx := int64(0)
	for i := 1; i < len(maxHeights); i++ {
		maxx = max(maxx, prefix[i-1]+suffix[i])
	}
	return maxx
}

func maxHeight(maxHeights []int) (int, []int) {
	height := -1
	index := []int{}
	for i := 0; i < len(maxHeights); i++ {
		if maxHeights[i] > height {
			height = maxHeights[i]
		}
	}

	for i, v := range maxHeights {
		if v == height {
			index = append(index, i)
		}
	}
	return height, index
}
