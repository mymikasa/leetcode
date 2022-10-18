package leetcode

func sumOfNumberAndReverse(num int) bool {

	for i := 0; i <= num; i++ {
		if i+reverseInt(i) == num {
			return true
		}
	}
	return false
}
