package main

type pair struct{ x, y int }

func judge(r int, c int, n int, par [][]int) bool {

	if c >= n || r >= n || c < 0 || r < 0 || par[r][c] > 0 {
		return false
	}

	return true
}

var dirs = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func generateMatrix(n int) [][]int {

	res := make([][]int, n)

	for i := range res {
		res[i] = make([]int, n)
	}

	row, col, dirIdx := 0, 0, 0
	dir := dirs[dirIdx]
	for i := 1; i <= n*n; i++ {

		//  当前位置合理
		if judge(row, col, n, res) {
			res[row][col] = i
		} else {
			i -= 1
			row -= dir.x
			col -= dir.y
			dirIdx = (dirIdx + 1) % 4
			dir = dirs[dirIdx]

		}
		row += dir.x
		col += dir.y

	}
	return res
}

func main() {
	// n := 3
	generateMatrix(3)
}
