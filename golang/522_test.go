package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLUSlength(t *testing.T) {
	strs := []string{"aa", "aaa", "aaa"}
	res := findLUSlength(strs)
	ans := -1
	assert.Equal(t, res, ans)
}
