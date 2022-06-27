package leetcode

func ConvertToTitle(columnNumber int) string {
	ans := []byte{}

	for columnNumber > 0 {
		temp := (columnNumber-1)%26 + 1
		ans = append(ans, 'A'+byte(temp-1))
		columnNumber = (columnNumber - temp) / 26
	}
	return string(ans)
}

//
