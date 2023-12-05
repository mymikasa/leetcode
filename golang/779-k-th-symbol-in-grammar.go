package leetcode

import "math"

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	mid := math.Pow(float64(2), float64(n-2))

	if k <= int(mid) {
		return kthGrammar(n-1, k)
	}
	return 1 ^ kthGrammar(n, k-int(mid))
}
