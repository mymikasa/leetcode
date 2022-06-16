package main

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)

	r, lenth := 0, len(nums)
	res := 0

	for l, num := range nums {
		if l == 0 || num != nums[l-1] {
			for r < lenth {
				if nums[r]-nums[l] < k {
					r++
				} else if nums[r]-nums[l] == k {
					r++
					res++
					break
				} else {
					break
				}
			}
		}
	}
	return res
}
