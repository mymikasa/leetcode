package leetcode

func HammingWeight(num uint32) int {
	ans := 0

	for num != 0 {
		if num&1 == 1 {
			ans++
		}

		num >>= 1
	}
	return ans

	// 给定一个数n，每进行一次n&(n-1)计算，其结果中都会少了一位1，而且是最后一位。
	// 可以通过不断地用n&(n-1)操作去掉n中最后一位1的方法求出n中1的个数

	// for num > 0 {
	// 	num &= num -1
	// 	count++
	// }
	// return count
}
