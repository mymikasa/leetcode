package main

import "math/bits"

//func maximumRows(matrix [][]int, numSelect int) int {
//	n := len(matrix)
//
//	mapp := make(map[int][]int)
//
//	for i, row := range matrix {
//		mapp[i] = []int{}
//		for j, _ := range row {
//			if matrix[i][j] == 1 {
//				mapp[i] = append(mapp[i], j)
//			}
//		}
//	}
//
//}

func maximumRows(matrix [][]int, numSelect int) (ans int) {
	m, n := len(matrix), len(matrix[0])
	mask := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mask[i] += matrix[i][j] << (n - j - 1)
		}
	}
	res, limit := 0, 1<<n
	for cur := 1; cur < limit; cur++ {
		if bits.OnesCount(uint(cur)) != numSelect {
			continue
		}
		t := 0
		for j := 0; j < m; j++ {
			if (mask[j] & cur) == mask[j] {
				t++
			}
		}
		res = max(res, t)
	}
	return res
}
