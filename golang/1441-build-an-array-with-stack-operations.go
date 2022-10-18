package leetcode

func buildArray(target []int, n int) []string {
	l := 0
	e := len(target)
	var res []string
	for i := 1; i <= n; i++ {
		if l >= e {
			break
		}
		if i == target[l] {
			res = append(res, "Push")
			l++
		} else if i < target[l] {
			res = append(res, "Push")
			res = append(res, "Pop")
		}
	}
	return res
}
