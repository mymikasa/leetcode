package leetcode

func countDistinctIntegers(nums []int) int {
	demo := make(map[int]bool)

	for _, i := range nums {
		demo[i] = true
		demo[reverseInt(i)] = true
	}
	return len(demo)
}

func reverseInt(num int) int {
	res := 0

	for num > 0 {
		temp := num % 10
		res = res*10 + temp
		num = num / 10
	}
	return res
}
