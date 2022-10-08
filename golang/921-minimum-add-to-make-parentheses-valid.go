package leetcode

import "strings"

func minAddToMakeValid(s string) int {
	for strings.Contains(s, "()") {
		s = strings.Replace(s, "()", "", -1)
	}
	return len(s)
}
