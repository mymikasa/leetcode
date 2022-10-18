package leetcode

func maxChunksToSorted(arr []int) int {

	res := 0
	var temp [][]int
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[i] {
				temp = append(temp, []int{i, j})
				flag = true
				break
			}
		}
		if !flag {
			temp = append(temp, []int{i, i})
		}
	}

	for i := 0; i < len(arr); i++ {
		if temp[i][0] != temp[i][1] {
			//	res++
			//} else {
			maxx := -1
			for j := i; j <= temp[i][1]; j++ {
				maxx = max(maxx, temp[j][1])
			}
			i = maxx
		}
		res++
	}
	return res
}
