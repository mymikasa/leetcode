package leetcode

import (
	"strings"
)

func stringToSlice(str string) []string {
	res := strings.Split(str, "")
	return res
}

func greaterEqualIntSlice(a, b []int) bool {
	if len(a) > len(b) {
		return true
	} else if len(a) < len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return false
		}
	}
	return true
}

func canTransform(start string, end string) bool {
	//输入: start = "RXXLRXRXL", end = "XRLXXRRLX"
	//输出: True
	r := len(start)

	ss := stringToSlice(start)
	se := stringToSlice(end)

	if strings.Replace(start, "X", "", -1) != strings.Replace(end, "X", "", -1) {
		return false
	}

	countSX := 0
	countEX := 0

	seL := []int{}
	ssL := []int{}
	for i := 0; i < r; i++ {
		if ss[i] == "X" {
			countSX++
		} else if ss[i] == "L" {
			ssL = append(ssL, countSX)
		}

		if se[i] == "X" {
			countEX++
		} else if se[i] == "L" {
			seL = append(seL, countEX)
		}
	}

	if !greaterEqualIntSlice(ssL, seL) {
		return false
	}

	countSX = 0
	countEX = 0

	seR := []int{}
	ssR := []int{}
	for i := r - 1; i >= 0; i-- {
		if ss[i] == "X" {
			countSX++
		} else if ss[i] == "R" {
			ssR = append(ssR, countSX)
		}

		if se[i] == "X" {
			countEX++
		} else if se[i] == "R" {
			seR = append(seR, countEX)
		}
	}
	if !greaterEqualIntSlice(ssR, seR) {
		return false
	}
	return true
}
