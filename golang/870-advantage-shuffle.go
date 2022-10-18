package leetcode

import "sort"

func binarySearchMaxMin(val int, pat []int, flag []bool) int {
	left := 0
	right := len(pat) - 1

	if val >= pat[right] {
		left = 0
	} else {
		for left < right {
			mid := (left + right) / 2
			if pat[mid] > val {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}

	for left < len(pat) && flag[left] {
		left++
	}

	// 没有找到合适的值， 直接取最小值
	if left >= len(pat) {
		left = 0
	}
	for left < len(pat) && flag[left] {
		left++
	}
	return left
}

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	idx1 := make([]bool, n)
	res := make([]int, n)

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	for i, v := range nums2 {
		idx := binarySearchMaxMin(v, nums1, idx1)
		res[i] = nums1[idx]
		idx1[idx] = true
	}
	return res
}
