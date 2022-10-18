package leetcode

func totalFruit(fruits []int) int {

	res := 0
	for i := 0; i < len(fruits); i++ {
		r := -1
		count := 1
		for j := i + 1; j < len(fruits); j++ {

			if fruits[j] != fruits[i] {

				if r == -1 {
					r = j
				} else {
					if fruits[r] == fruits[j] {
						count++
						r = j
						continue
					}
					i = r - 1
					res = max(count, res)
					break
				}
			}
			count++
		}
		res = max(count, res)
		for j := r; j >= 0; j-- {
			if fruits[j] != fruits[r] {
				i = j
				break
			}
		}
	}
	return res
}
