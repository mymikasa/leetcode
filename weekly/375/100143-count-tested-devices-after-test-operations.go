package main

import (
	"slices"
)

func countTestedDevices(batteryPercentages []int) int {

	cnt := 0
	res := 0
	for _, v := range batteryPercentages {
		if v-cnt > 0 {
			res++
			cnt++
		}
	}
	return res
}

func countSubarrays(nums []int, k int) int64 {
	//	 [1,3,2,3,3], k = 2
	//1 2 3 4  12 2  4  124 24
	// 1323 323 13233 3233 233 33
	//l := len(nums)
	mx := slices.Max(nums)
	cntMx, left := 0, 0
	var ans int64
	for _, x := range nums {
		if x == mx {
			cntMx++
		}
		for cntMx == k {
			if nums[left] == mx {
				cntMx--
			}
			left++
		}
		ans += int64(left)
	}
	return ans

}
