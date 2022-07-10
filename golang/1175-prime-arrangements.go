package leetcode

import "math"

func primeNums(n int) int {
	temp := make([]bool, n+1)
	temp[0], temp[1] = true, true
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for j := i; i*j <= n; j++ {
			temp[i*j] = true
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		if !temp[i] {
			res++
		}
	}
	return res
}

func factorial(n int) int {
	ret := 1
	if n == 0 {
		return 1
	}
	for i := 1; i <= n; i++ {
		ret *= i
		ret = ret % int((math.Pow(10, 9) + 7))
	}
	return ret
}

func numPrimeArrangements(n int) int {

	pn := primeNums(n)
	return factorial(pn) * factorial(n-pn) % int((math.Pow(10, 9) + 7))
}
