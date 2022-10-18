package leetcode

//	func abs(num int) int {
//		if num < 0{
//			return
//		}
//	}
func findMaxK(nums []int) int {
	//sort.Ints(nums)
	demo := make(map[int]bool)

	for i := -1001; i <= 1001; i++ {
		demo[i] = false
	}

	for _, i := range nums {
		demo[i] = true
	}

	res := -1
	for _, i := range nums {
		if demo[i] == true && demo[-i] == true {
			res = max(res, abs(i))
		}
	}
	return res
}
