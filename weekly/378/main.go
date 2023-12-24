package main

import (
	"sort"
)

func numberGame(nums []int) []int {

	var arr []int
	sort.Ints(nums)

	for i := 0; i < len(nums); i += 2 {
		arr = append(arr, nums[i+1], nums[i])
	}

	return arr
}

func containsMapInt(mp map[int]bool, element int) bool {
	_, ok := mp[element]
	return ok
}

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {

	hFences = append(hFences, 1, m)
	vFences = append(vFences, 1, n)
	sort.Ints(hFences)
	sort.Ints(vFences)

	hn := len(hFences)
	hv := len(vFences)

	var arrH []int
	var arrV []int

	for i := 0; i < hn; i++ {
		for j := i + 1; j < hn; j++ {
			arrH = append(arrH, hFences[j]-hFences[i])
		}
	}

	for i := 0; i < hv; i++ {
		for j := i + 1; j < hv; j++ {
			arrV = append(arrV, vFences[j]-vFences[i])
		}
	}

	sort.Ints(arrV)

	mapInt := map[int]bool{}

	for _, v := range arrH {
		mapInt[v] = true
	}

	for i := len(arrV) - 1; i >= 0; i-- {
		if containsMapInt(mapInt, arrV[i]) {
			return arrV[i] * arrV[i] % (1e9 + 7)
		}

	}
	return -1
}

func minimumCost(source string, target string, original, changed []byte, cost []int) (ans int64) {
	dis := [26][26]int{}

	for i := range dis {
		for j := range dis[i] {
			if j != i {
				dis[i][j] = 1e13
			}
		}
	}

	for i, v := range cost {
		x := original[i] - 'a'
		y := changed[i] - 'a'

		dis[x][y] = min(dis[x][y], v)
	}

	for k := range dis {
		for i := range dis {
			for j := range dis {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	for i, v := range source {
		ans += int64(dis[v-'a'][target[i]-'a'])
	}
	if ans >= 1e13 {
		return -1
	}
	return ans
}
