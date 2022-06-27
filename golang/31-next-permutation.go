package leetcode

func Reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func NextPermutation(nums []int) {

	l := len(nums)
	change := -1
	ind := -1

	for i := 0; i < l; i++ {
		temp := 1000
		for j := i + 1; j < l; j++ {
			if nums[i] < nums[j] && nums[j] < temp {
				ind = i
				change = j
				temp = nums[j]
				// nums[i], nums[j] = nums[j], nums[i]
			}
		}

	}
	if change == -1 {
		Reverse(nums)
	} else {
		nums[ind], nums[change] = nums[change], nums[ind]

		for i := ind + 1; i < l; i++ {
			temp := 1000
			for j := i; j < l; j++ {
				if nums[i] > nums[j] && nums[j] < temp {
					nums[i], nums[j] = nums[j], nums[i]
					temp = nums[j]
				}
			}
		}
	}
}
