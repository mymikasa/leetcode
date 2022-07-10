package leetcode

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	var res [][]int

	maxx := 10000000
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < maxx {
			// clear slice
			res = [][]int{{arr[i-1], arr[i]}}
		} else if arr[i]-arr[i-1] == maxx {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
