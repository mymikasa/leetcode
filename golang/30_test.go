package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	res := findSubstring(s, words)
	ans := []int{0, 9}
	assert.Equal(t, res, ans)

}
