package main

func incremovableSubarrayCount(nums []int) int {

	n := len(nums)

	ind := 1
	ans := 0
	for i := 0; i < n-1; i++ {

		j := i + 1

		if nums[j] > nums[i] {
			ind++
		} else {
			break
		}

	}
	ind2 := 1
	for i := n - 1; i > 0; i-- {
		j := i - 1

		if nums[j] < nums[i] {
			ind2++
		} else {
			break
		}
	}

	//less := false
	if ind2 == ind && ind == n {
		return n * (n + 1) / 2
	}

	ans += ind2
	ans += ind
	ans += 1

	temp := n

	for i := ind; i < n; i++ {
		if nums[i] > nums[ind] {
			temp = i
			break
		}
	}

	ans += n - temp
	return ans
}

//[6,5,7,8]

// 6
// 578  6 65 657 6578
//[5], [6], [5,7], , , ] 和
