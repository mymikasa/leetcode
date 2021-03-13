package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c && abs(arr[i]-arr[j]) <= a {
					count += 1
				}
			}
		}
	}

	return count
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return 0 - a
}

// func main() {
// 	arr := []int{3, 0, 1, 1, 9, 7}
// 	a := 7
// 	b := 3
// 	c := 2

// 	fmt.Println(countGoodTriplets(arr, a, b, c))

// }
