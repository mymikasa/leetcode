package leetcode

import (
	"sort"
	"testing"
)

func TestReplaceWords(t *testing.T) {
	aa := []string{"aa", "bbbb", "c"}

	sort.Slice(aa, func(i, j int) bool {
		return len(aa[i]) < len(aa[j])
	})
	t.Log(aa)
	t.Log(aa[2][:2])
}
