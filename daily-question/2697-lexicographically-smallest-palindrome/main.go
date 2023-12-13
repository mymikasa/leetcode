package main

func makeSmallestPalindrome(s string) string {
	res := []byte{}

	j := len(s)
	for i := 0; i < len(s); i++ {
		j -= 1
		minn := min(s[i], s[j])
		res = append(res, minn)
	}

	return string(res)
}
