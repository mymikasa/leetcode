package leetcode

func countSubarrays(nums []int, minK int, maxK int) (res int64) {

	sumK := minK + maxK
	start := 0
	var count int64

	for i := 0; i < len(nums); i++ {
		j := nums[start]
		if j == minK || j == maxK {
			start = i
		} else {
			continue
		}
		end := -1
		for k := start + 1; k < len(nums); k++ {
			if nums[k] == sumK-j {
				end = k
				break
			}
		}
		if end == -1 {
			break
		} else {
			count = 0
			for i = end; i < len(nums); i++ {
				if nums[i] >= minK && nums[i] <= maxK {
					count++
				} else {
					break
				}
			}

			if minK == maxK {
				count++
				res += count * (count + 1) / 2
			} else {
				res += count*(count+1)/2 - 1
			}
		}
	}
	return res
}
