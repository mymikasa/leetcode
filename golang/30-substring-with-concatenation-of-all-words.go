package leetcode

// timeout

// func findSubstring(s string, words []string) []int {

// 	ans := []int{}
// 	ls := len(s)

// 	lws := len(words)

// 	lw := len(words[0])

// 	for i := 0; i < lw; i++ {
// 		ss := map[string]int{}
// 		for _, v := range words {
// 			ss[v] += 1
// 		}

// 		for j := i; j < ls; j += lw {
// 			tt := ""
// 			if j+lw > ls {
// 				break
// 			}
// 			for k := j; k < j+lw; k++ {
// 				tt = tt + string(s[k])
// 			}

// 			if _, ok := ss[tt]; ok {
// 				ss[tt]--
// 				if ss[tt] < 0 {
// 					break
// 				}
// 			} else {
// 				break
// 			}

// 			if j-i+lw == lws*lw {
// 				ans = append(ans, j)
// 			}
// 		}
// 	}
// 	return ans
// }

func findSubstring(s string, words []string) []int {

	ans := []int{}
	ls := len(s)

	lws := len(words)

	lw := len(words[0])

	for i := 0; i < lw; i++ {
		ss := map[string]int{}
		for _, v := range words {
			ss[v] += 1
		}

		for j := i; j < ls; j += lw {
			tt := ""
			if j+lw > ls {
				break
			}
			for k := j; k < j+lw; k++ {
				tt = tt + string(s[k])
			}

			if _, ok := ss[tt]; ok {
				ss[tt]--
				if ss[tt] < 0 {
					break
				}
			} else {
				break
			}

			if j-i+lw == lws*lw {
				ans = append(ans, j)
			}
		}
	}
	return ans
}
