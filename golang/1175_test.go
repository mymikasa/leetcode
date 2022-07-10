package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumPrimeArrangements(t *testing.T) {
	n := 100
	ans := 682289015

	res := numPrimeArrangements(n)
	assert.Equal(t, res, ans)
}
