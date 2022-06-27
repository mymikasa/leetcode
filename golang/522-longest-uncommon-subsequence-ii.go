package leetcode

func isSubStr(s, t string) bool {
	ptS := 0
	for ptT := range t {
		if s[ptS] == t[ptT] {
			ptS++
			if ptS == len(s) {
				return true
			}
		}
	}
	return false
}

func findLUSlength(strs []string) int {
	ans := -1

	ls := len(strs)
	for i := 0; i < ls; i++ {
		demo := 0
		for j := 0; j < ls; j++ {
			if isSubStr(strs[i], strs[j]) || i == j {
				demo++
				continue
			}
		}
		if len(strs[i]) > ans && demo == 1 {
			ans = len(strs[i])
		}
	}
	return ans
}
