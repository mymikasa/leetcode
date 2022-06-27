package leetcode

func Generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1

		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}

	return res
}

// func getRow(rowIndex int) []int {

// 	res := make([][]int, rowIndex)

// 	for i := range res {
// 		res[i] = make([]int, i+1)
// 		res[i][0] = 1
// 		res[i][i] = 1

// 		for j := 1; j < i; j++ {
// 			res[i][j] = res[i-1][j] + res[i-1][j-1]
// 		}
// 	}

// 	return res[rowIndex]

// }

func GetRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}
