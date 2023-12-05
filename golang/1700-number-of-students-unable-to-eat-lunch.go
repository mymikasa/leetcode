package leetcode

func countStudents(students []int, sandwiches []int) int {
	s0 := 0
	for i := 0; i < len(students); i++ {
		if students[i] == 0 {
			s0++
		}
	}
	s1 := len(students) - s0
	if s1 == s0 {
		return 0
	}
	for i := 0; i < len(students); i++ {
		if sandwiches[i] == 0 {
			s0--
		} else {
			s1--
		}

		if s0 < 0 || s1 < 0 {
			break
		}
	}
	return s0 + s1 + 1
}
