package main

func minReorder(n int, connections [][]int) int {

	e := make(map[int][][]int)
	known := make(map[int]bool)
	for i := 0; i < n; i++ {
		known[i] = false
	}
	// 建图
	for _, v := range connections {
		e[v[0]] = append(e[v[0]], []int{v[1], 1})
		e[v[1]] = append(e[v[1]], []int{v[0], 0})
	}

	var dfs func(int)
	res := 0
	dfs = func(start int) {
		if known[start] {
			return
		}
		known[start] = true
		for i, v := range e[start] {
			if !known[v[0]] {
				dfs(v[0])
			} else {
				if v[1] == 0 {
					res += 1
					e[start][i][1] = 0
				}
			}
		}
	}
	dfs(0)
	return res
}
