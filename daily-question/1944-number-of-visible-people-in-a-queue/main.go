package main

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	res := make([]int, n)
	st := make([]int, 0)

	for i := n - 1; i >= 0; i-- {

		h := heights[i]
		//10, 6, 8, 5, 11, 9
		for len(st) > 0 && h >= st[len(st)-1] {
			st = st[:len(st)-1]
			res[i] += 1
		}

		if len(st) > 0 {
			res[i] += 1
		}
		st = append(st, h)
	}

	return res
}
