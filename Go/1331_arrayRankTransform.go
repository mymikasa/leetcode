package main

import "sort"

func arrayRankTransform(arr []int) []int {

	temp := make([]int, len(arr))
	copy(temp, arr)
	start := 1

	resMap := make(map[int]int)

	sort.Ints(arr)
	for _, v := range arr {

		if _, ok := resMap[v]; ok {
			continue
		}
		resMap[v] = start
		start++
	}

	for i, v := range temp {
		temp[i] = resMap[v]
	}
	return temp
}
