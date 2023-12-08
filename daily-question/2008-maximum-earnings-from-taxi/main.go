package main

import (
	"sort"
)

func maxTaxiEarnings(n int, rides [][]int) int64 {
	// n = 20, rides = [[1,6,11],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
	dp := make([]int64, len(rides)+1)
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})
	var max func(a, b int64) int64
	max = func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	dp[0] = int64(rides[0][1] - rides[0][0] + rides[0][2])

	for i := 1; i < len(rides); i++ {
		j := sort.Search(i, func(k int) bool {
			return rides[k][1] > rides[i][0]
		})

		j -= 1
		if j == i || j == -1 {
			dp[i] = max(dp[i-1], int64(rides[i][1]-rides[i][0]+rides[i][2]))
		} else {
			dp[i] = max(dp[i-1], int64(rides[i][1]-rides[i][0]+rides[i][2])+dp[j])
		}
	}

	return dp[len(rides)-1]
}
