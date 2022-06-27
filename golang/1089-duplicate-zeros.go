package leetcode

//[1,0,2,3,0,4,5,0]
//[1,0,0,2,3,0,0,4]

func DuplicateZeros(arr []int) {
	r := 0
	lenth := len(arr)
	ar := 0

	for i, v := range arr {
		r++
		if v == 0 {
			r++
		}
		if r >= lenth {
			ar = i
			break
		}
	}

	// 最后一个值为0，且不能重复
	if r > lenth {
		r -= 2
		arr[r] = arr[ar]
		r--
	} else {
		r--
		if arr[ar] == 0 {
			arr[r] = arr[ar]
			arr[r-1] = arr[ar]
			r -= 2
		} else {
			arr[r] = arr[ar]
			r -= 1
		}
	}

	// if r > lenth {
	// 	r--
	// }

	// if r == lenth {
	// 	r--
	// 	arr[r] = arr[ar]
	// 	r--
	// }

	for j := ar - 1; j >= 0; j-- {
		if arr[j] == 0 {
			arr[r] = 0
			r--
		}
		arr[r] = arr[j]
		r--
	}
}
