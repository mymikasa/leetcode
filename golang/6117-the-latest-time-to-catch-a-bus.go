package leetcode

import "sort"

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	// [20,30,10]
	// [19,13,26,4,25,11,21]

	val := []int{}
	mapp := make(map[int]int)
	ss := 0
	for _, v := range buses {
		done := 0
		for ss < len(passengers) {
			j := passengers[ss]
			if j <= v {
				val = append(val, j)
				done++
				mapp[v] = done
				ss++
			}
			if done == capacity {
				break
			}
			if j > v {
				break
			}

		}

	}

	// 不需要插队
	if mapp[buses[len(buses)-1]] < capacity {
		res := buses[len(buses)-1]
		for i := len(val) - 1; i >= 0; i-- {
			if res == val[i] {
				res--
			} else {
				return res
			}
		}
		return res
	}

	ll := len(val)
	if ll == 1 {
		return val[0] - 1
	}
	nex := val[len(val)-1] - 1

	for i := len(val) - 2; i >= 0; i-- {
		if nex > val[i] {
			return nex
		} else {
			nex--
		}
	}

	return nex
}
