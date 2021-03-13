package main

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	sort.Ints(append(nums1[:m], nums2...))
}

// func main() {
// 	fmt.Println("hh")
// }
