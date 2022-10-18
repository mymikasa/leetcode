package leetcode

func distinctSubseqII(s string) int {
	const mod int = 1e9 + 7
	g := make([]int, 26)

	for _, i := range s {
		total := 1

		for _, j := range g {
			total = (total + j) % mod
		}

		g[i-'a'] = total
	}

	res := 0
	for _, i := range g {
		res = (res + i) % mod
	}
	return res
}
