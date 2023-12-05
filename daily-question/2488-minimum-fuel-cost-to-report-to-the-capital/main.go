package main

func minimumFuelCost(roads [][]int, seats int) int64 {
	// roads = [[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]], seats = 2
	g := make([][]int, len(roads)+1)

	// 联通图
	for _, v := range roads {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}

	var res int64

	var dfs func(int, int) int

	dfs = func(start, end int) int {
		peopleSum := 1
		for _, v := range g[start] {
			if v != end {
				peopleCnt := dfs(v, start)
				peopleSum += peopleCnt
				res = res + int64((peopleCnt+seats-1)/seats)
			}
		}
		return peopleSum
	}
	dfs(0, -1)
	return res
}

func main() {
	// var roads [][]int
	// 3,1],[3,2],[1,0],[0,4],[0,5],[4,6
	roads := [][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}
	println(minimumFuelCost(roads, 2))
}
