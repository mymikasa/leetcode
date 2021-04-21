package main

func twoSum(numbers []int, target int) []int {
	lenth := len(numbers)

	for i := 0; i < lenth; i++ {
		left, right := i+1, lenth-1

		for left <= right {
			mid := (right-left)/2 + left

			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return []int{-1, -1}
}
