package main

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	prefix := getLISArray(nums)
	suffix := getLISArray(reversed(nums))
	suffix = reversed(suffix)
	//2,1,1,5,6,2,3,1
	res := n
	for i := 0; i < n; i++ {

		res = min(res, n-prefix[i]-suffix[i]+1)

	}
	return res
}

func getLISArray(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	return dp
}

func reversed(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = nums[n-i-1]
	}
	return res
}
